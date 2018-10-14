package peg

import (
	"fmt"
	"testing"

	"github.com/lioneagle/abnf/src/charset"

	"github.com/lioneagle/goutil/src/test"
)

func TestPegAndPredicate(t *testing.T) {
	testdata := []struct {
		chars       string
		exprName    string
		closureName string
		any         charset.Range
		inverse     bool
		str         string
	}{
		{"a", "", "", charset.Range{}, false, "&('a')"},
		{"a-bA-F-", "", "", charset.Range{}, false, "&(['-', 'A'-'F', 'a'-'b'])"},
		{"a-b", "", "r1", charset.Range{Low: 'a', High: 'f'}, true, "r1"},
		{"a-b", "test", "", charset.Range{Low: 'a', High: 'f'}, true, "&test"},
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
			expr2 := NewPegAndPredicate(v.closureName, expr1)
			str := expr2.String()

			test.EXPECT_EQ(t, str, v.str, "")
		})
	}
}
