package peg

import (
	"fmt"
	"testing"

	"github.com/lioneagle/abnf/src/charset"

	"github.com/lioneagle/goutil/src/test"
)

func TestPegClosure(t *testing.T) {
	testdata := []struct {
		chars       string
		exprName    string
		closureName string
		any         charset.Range
		inverse     bool
		min         int
		max         int
		str         string
	}{
		{"a", "", "", charset.Range{}, false, 1, PEG_INFINITE_NUM, "1*('a')"},
		{"a-bA-F-", "", "", charset.Range{}, false, 0, PEG_INFINITE_NUM, "*(['-', 'A'-'F', 'a'-'b'])"},
		{"a-b", "", "r1", charset.Range{Low: 'a', High: 'f'}, true, 3, 7, "r1"},
		{"a-b", "test", "", charset.Range{Low: 'a', High: 'f'}, true, 0, 7, "*7test"},
		{"a-b", "test", "", charset.Range{Low: 'a', High: 'f'}, true, 0, 1, "[test]"},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			expr1 := NewPegCharset(v.exprName, nil)
			if !v.inverse {
				expr1.MakeFromBytes([]byte(v.chars))
			} else {
				expr1.MakeFromBytesInverse(&v.any, []byte(v.chars))
			}
			expr2 := NewPegClosure(v.closureName, expr1, v.min, v.max)
			str := expr2.String()

			test.EXPECT_EQ(t, str, v.str, "")
		})
	}
}
