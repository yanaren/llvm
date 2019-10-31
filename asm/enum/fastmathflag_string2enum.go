// Code generated by "string2enum -linecomment -type FastMathFlag ../../ir/enum"; DO NOT EDIT.

package enum

import (
	"fmt"

	"github.com/llir/llvm/ir/enum"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the string2enum command to generate them again.
	var x [1]struct{}
	_ = x[enum.FastMathFlagAFn-0]
	_ = x[enum.FastMathFlagARcp-1]
	_ = x[enum.FastMathFlagContract-2]
	_ = x[enum.FastMathFlagFast-3]
	_ = x[enum.FastMathFlagNInf-4]
	_ = x[enum.FastMathFlagNNaN-5]
	_ = x[enum.FastMathFlagNSZ-6]
	_ = x[enum.FastMathFlagReassoc-7]
}

const _FastMathFlag_name = "afnarcpcontractfastninfnnannszreassoc"

var _FastMathFlag_index = [...]uint8{0, 3, 7, 15, 19, 23, 27, 30, 37}

// FastMathFlagFromString returns the FastMathFlag enum corresponding to s.
func FastMathFlagFromString(s string) enum.FastMathFlag {
	if len(s) == 0 {
		return 0
	}
	for i := range _FastMathFlag_index[:len(_FastMathFlag_index)-1] {
		if s == _FastMathFlag_name[_FastMathFlag_index[i]:_FastMathFlag_index[i+1]] {
			return enum.FastMathFlag(i)
		}
	}
	panic(fmt.Errorf("unable to locate FastMathFlag enum corresponding to %q", s))
}