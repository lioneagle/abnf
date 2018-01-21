package charset

import (
	"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

type charsetExprTestNode struct {
	name      string
	wellKnown bool
	charsets  []string
	op        uint32
}

func (this *charsetExprTestNode) buildCharset() *Charset {
	c := NewCharset()
	c.Name = this.name
	c.IsWellKnown = this.wellKnown
	for _, v := range this.charsets {
		c2 := NewCharset()
		c2.MakeFromBytes([]byte(v))
		c.UniteCharset(c2)
	}
	return c
}

func buildCharsetExpr(nodes []charsetExprTestNode) *CharsetExpr {
	expr := &CharsetExpr{}
	for _, v := range nodes {
		c := v.buildCharset()
		if v.op == CHARSET_OP_PLUS {
			expr.Plus(c)
		} else if v.op == CHARSET_OP_MINUS {
			expr.Minus(c)
		}
	}
	return expr
}

func TestCharsetExpr(t *testing.T) {
	type result struct {
		str                 string
		size                uint32
		charsetNum          uint32
		empty               bool
		plusNum             uint32
		minusNum            uint32
		branchNum           uint32
		hasWellKnownCharset bool
	}

	testdata := []struct {
		expr   []charsetExprTestNode
		wanted *result
	}{
		{[]charsetExprTestNode{}, &result{"", 0, 0, true, 0, 0, 0, false}},
		{[]charsetExprTestNode{{"alpha", true, []string{"a-z", "A-Z"}, CHARSET_OP_PLUS}}, &result{"{alpha}", 52, 1, false, 1, 0, 1, true}},
		{[]charsetExprTestNode{{"alpha", true, []string{"a-z", "A-Z"}, CHARSET_OP_MINUS}}, &result{"", 0, 0, true, 0, 0, 0, false}},
		{[]charsetExprTestNode{{"alpha", true, []string{"a-z", "A-Z"}, CHARSET_OP_PLUS}, {"digit", true, []string{"0-9"}, CHARSET_OP_PLUS}}, &result{"{alpha} + {digit}", 62, 2, false, 2, 0, 2, true}},
		{[]charsetExprTestNode{{"alpha", true, []string{"a-z", "A-Z"}, CHARSET_OP_PLUS}, {"digit", true, []string{"0-9"}, CHARSET_OP_MINUS}}, &result{"{alpha} - {digit}", 52, 2, false, 1, 1, 2, true}},
		{[]charsetExprTestNode{{"alpha", true, []string{"a-z", "A-Z"}, CHARSET_OP_PLUS}, {"token", false, []string{"a-b"}, CHARSET_OP_MINUS}}, &result{"{alpha} - {a, b}", 50, 2, false, 1, 1, 3, true}},
		{[]charsetExprTestNode{{"alpha", true, []string{"a-z", "A-Z"}, CHARSET_OP_MINUS}, {"token", false, []string{"a-b"}, CHARSET_OP_PLUS}}, &result{"{a, b}", 2, 1, false, 1, 0, 2, false}},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			expr := buildCharsetExpr(v.expr)
			ret := &result{}
			ret.str = expr.String()
			ret.size = expr.CharsetSize()
			ret.charsetNum = expr.CharsetNum()
			ret.empty = expr.Empty()
			ret.plusNum = expr.PlusNum()
			ret.minusNum = expr.MinusNum()
			ret.branchNum = expr.BranchNum()
			ret.hasWellKnownCharset = expr.HasWellKnownCharset()

			ok, msg := test.DiffEx("", ret, v.wanted)
			if !ok {
				t.Errorf("\n" + msg)
			}
		})
	}
}
