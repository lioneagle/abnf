package regexpr

import (
	//"bytes"
	//"os"
	//"strconv"
	"testing"
	"trace"
)

func TestRegExprCat(t *testing.T) {
	testdata := []struct {
		name  string
		expr1 RegExpr
		expr2 RegExpr
		str   string
	}{
		{"", NewRegExprCharset("", []byte("a")), NewRegExprCharset("", []byte("A-G")), "\"a\" [A-G]"},
		{"", NewRegExprCharset("xyz", []byte("a")), NewRegExprCharset("", []byte("A-G")), "xyz [A-G]"},
		{"abc", NewRegExprCharset("xyz", []byte("a")), NewRegExprCharset("", []byte("A-G")), "abc"},
	}
	prefix := trace.CallerName(0)

	for i, v := range testdata {
		cat := NewRegExprCat(v.name, v.expr1, v.expr2)
		str := cat.String()

		if str != v.str {
			t.Errorf("%s[%d] failed: str = %s, wanted = %s\n", prefix, i, str, v.str)
		}
	}
}
