package regexpr

import (
	"fmt"
	"testing"

	"github.com/lioneagle/abnf/src/charset"

	"github.com/lioneagle/goutil/src/test"
)

func TestRegExprCharset(t *testing.T) {
	testdata := []struct {
		c       string
		name    string
		any     charset.Range
		inverse bool
		str     string
	}{
		{"", "", charset.Range{}, false, ""},
		{"a", "", charset.Range{}, false, "'a'"},
		{"a-bA-F-", "", charset.Range{}, false, "['-', 'A'-'F', 'a'-'b']"},
		{"a-b", "", charset.Range{Low: 'a', High: 'f'}, true, "['c'-'e']"},
		{"a-b", "test", charset.Range{Low: 'a', High: 'f'}, true, "test"},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			expr := NewRegExprCharset(v.name, nil)
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
