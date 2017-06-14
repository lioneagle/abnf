package regexpr

import (
	//"bytes"
	//"os"
	//"strconv"
	"testing"
	"trace"
)

func TestRegExprAlt(t *testing.T) {
	testdata := []struct {
		name string
		expr RegExpr
		str  string
	}{
		{"", NewRegExprCharset("", []byte("a")), "(\"a\")?"},
		{"", NewRegExprCharset("xyz", []byte("a")), "xyz?"},
		{"abc", NewRegExprCharset("xyz", []byte("a")), "abc"},
		{"", NewRegExprCharset("alpha", []byte("a-zA-Z")), "alpha?"},
	}
	prefix := trace.CallerName(0)

	for i, v := range testdata {
		cat := NewRegExprAlt(v.name, v.expr)
		str := cat.String()

		if str != v.str {
			t.Errorf("%s[%d] failed: str = %s, wanted = %s\n", prefix, i, str, v.str)
		}
	}
}
