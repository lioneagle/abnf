package regexpr

import (
	"charset"
	//"bytes"
	//"os"
	//"strconv"
	"testing"
	"trace"
)

func TestRegExprCharset(t *testing.T) {
	testdata := []struct {
		c       string
		name    string
		any     charset.Range
		inverse bool
		str     string
	}{
		{"a-bA-F-", "", charset.Range{}, false, "[\\-, A-F, a]"},
		{"a-b", "", charset.Range{'a', 'f'}, true, "[b-f]"},
		{"a-b", "test", charset.Range{'a', 'f'}, true, "test"},
	}
	prefix := trace.CallerName(0)

	for i, v := range testdata {
		expr := NewRegExprCharset()
		expr.Name = v.name
		if !v.inverse {
			expr.MakeFromBytes([]byte(v.c))
		} else {
			expr.MakeFromBytesInverse(&v.any, []byte(v.c))
		}
		str := expr.String()

		if str != v.str {
			t.Errorf("%s[%d] failed: str = %s, wanted = %s\n", prefix, i, str, v.str)
		}
	}
}
