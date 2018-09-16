package peg

import (
	"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestPegKeys(t *testing.T) {
	testdata := []struct {
		name                 string
		expr1                string
		expr2                string
		isExpr1CaseSensitive bool
		isExpr2CaseSensitive bool
		str                  string
	}{
		{"", "test", "abc", true, true, `('test' / 'abc')`},
		{"", "test", "abc", false, true, `("test" / 'abc')`},
		{"uri-char", "test", "abc", false, false, `uri-char`},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			expr := NewPegKeys(v.name)
			expr.AppendKey(NewPegConstString("", []byte(v.expr1), v.isExpr1CaseSensitive))
			expr.AppendKey(NewPegConstString("", []byte(v.expr2), v.isExpr2CaseSensitive))
			str := expr.String()

			test.EXPECT_EQ(t, str, v.str, "")
		})
	}
}
