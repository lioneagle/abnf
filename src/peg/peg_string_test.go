package peg

import (
	"fmt"
	"testing"

	"github.com/lioneagle/abnf/src/charset"

	"github.com/lioneagle/goutil/src/test"
)

func TestPegString(t *testing.T) {
	testdata := []struct {
		c       string
		name    string
		any     charset.Range
		inverse bool
		min     int
		max     int
		str     string
	}{
		{"", "", charset.Range{}, false, 1, 5, ""},
		{"a", "", charset.Range{}, false, 1, PEG_INFINITE_NUM, "1*'a'"},
		{"a-bA-F-", "", charset.Range{}, false, 0, PEG_INFINITE_NUM, "*['-', 'A'-'F', 'a'-'b']"},
		{"a-b", "", charset.Range{Low: 'a', High: 'f'}, true, 3, 7, "3*7['c'-'e']"},
		{"a-b", "test", charset.Range{Low: 'a', High: 'f'}, true, 0, PEG_INFINITE_NUM, "test"},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			expr := NewPegString(v.name, nil, v.min, v.max)
			if !v.inverse {
				expr.MakeFromBytes([]byte(v.c))
			} else {
				expr.MakeFromBytesInverse(&v.any, []byte(v.c))
			}
			str := expr.String()

			test.EXPECT_EQ(t, str, v.str, "")
		})
	}
}
