package constant_test

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"

	"github.com/llir/llvm/asm"
	"github.com/mewkiz/pkg/diffutil"
	"github.com/mewkiz/pkg/osutil"
)

// words specifies whether to colour words in diff output.
const words = false

func TestModule(t *testing.T) {
	golden := []struct {
		path string
	}{
		// LLVM IR compatibility.
		{path: "../../testdata/llvm/test/Bitcode/compatibility.ll"},
		// Hex floating-point constants.
		{path: "../../testdata/llvm/test/Assembler/2002-04-07-HexFloatConstants.ll"},
		{path: "../../testdata/llvm/test/Assembler/half-constprop.ll"},
		// Constant expressions.
		{path: "../../testdata/llvm/test/Transforms/ConstProp/constant-expr.ll"},
		{path: "../../testdata/llvm/test/Assembler/insertextractvalue.ll"},
		{path: "../../testdata/llvm/test/DebugInfo/ARM/selectiondag-deadcode.ll"},
		{path: "../../testdata/llvm/test/Transforms/InstCombine/fma.ll"},
		{path: "../../testdata/llvm/test/Transforms/InstCombine/vec_demanded_elts.ll"},
		{path: "../../testdata/llvm/test/Transforms/InstCombine/vector_insertelt_shuffle.ll"},
		// Coreutils.
		{path: "../../testdata/coreutils/test/timeout.ll"},
		{path: "../../testdata/coreutils/test/vdir.ll"},
	}
	hasTestdata := osutil.Exists("../../testdata/llvm")
	for _, g := range golden {
		if filepath.HasPrefix(g.path, "../../testdata") && !hasTestdata {
			// Skip test cases from the llir/testdata submodule if not downloaded.
			// Users may add this submodule using git clone --recursive.
			continue
		}
		log.Printf("=== [ %s ] ===", g.path)
		m, err := asm.ParseFile(g.path)
		if err != nil {
			t.Errorf("unable to parse %q into AST; %+v", g.path, err)
			continue
		}
		path := g.path
		if osutil.Exists(g.path + ".golden") {
			path = g.path + ".golden"
		}
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			t.Errorf("unable to read %q; %+v", path, err)
			continue
		}
		want := string(buf)
		got := m.String()
		if want != got {
			if err := diffutil.Diff(want, got, words, filepath.Base(path)); err != nil {
				panic(err)
			}
			t.Errorf("module mismatch %q; expected `%s`, got `%s`", path, want, got)
			continue
		}
	}
}
