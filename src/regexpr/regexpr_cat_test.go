package regexpr

import (
	"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestRegExprCat(t *testing.T) {
	testdata := []struct {
		name  string
		expr1 RegExpr
		expr2 RegExpr
		str   string
	}{
		{"", NewRegExprCharset("", []byte("a")), NewRegExprCharset("", []byte("A-G")), "'a' ['A'-'G']"},
		{"", NewRegExprCharset("xyz", []byte("a")), NewRegExprCharset("", []byte("A-G")), "xyz ['A'-'G']"},
		{"abc", NewRegExprCharset("xyz", []byte("a")), NewRegExprCharset("", []byte("A-G")), "abc"},
		{"", NewRegExprCharset("alpha", []byte("a-zA-Z")), NewRegExprCharset("digit", []byte("0-9")), "alpha digit"},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			cat := NewRegExprCat(v.name, v.expr1, v.expr2)
			str := cat.String()

			test.EXPECT_EQ(t, str, v.str, "")
		})
	}
}
