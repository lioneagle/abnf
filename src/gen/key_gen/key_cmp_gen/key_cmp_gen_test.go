package key_cmp_gen

import (
	"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestBranchEqualNoCase(t *testing.T) {
	testdata := []struct {
		lhs byte
		rhs byte
		ret bool
	}{
		{'1', '1', true},
		{'1', '2', false},
		{'a', '2', false},
		{'1', 'z', false},
		{'a', 'a', true},
		{'a', 'A', true},
		{'z', 'Z', true},
	}

	for i, v := range testdata {
		v := v
		branch := &Branch{Value: v.lhs, Next: nil}
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			test.EXPECT_EQ(t, branch.EqualNoCase(v.rhs), v.ret, "")
		})
	}

}
