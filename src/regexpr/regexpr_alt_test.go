package regexpr

import (
	"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestRegExprAlt(t *testing.T) {
	testdata := []struct {
		name string
		expr RegExpr
		str  string
	}{
		{"", NewRegExprCharset("", []byte("a")), "('a')?"},
		{"", NewRegExprCharset("xyz", []byte("a")), "xyz?"},
		{"abc", NewRegExprCharset("xyz", []byte("a")), "abc"},
		{"", NewRegExprCharset("alpha", []byte("a-zA-Z")), "alpha?"},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			cat := NewRegExprAlt(v.name, v.expr)
			str := cat.String()
			test.EXPECT_EQ(t, str, v.str, "")
		})
	}
}
