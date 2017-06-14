package charset

import (
	//"bytes"
	//"fmt"
	//"os"
	//"strconv"
	"testing"
	"trace"
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
	testdata := []struct {
		expr         []charsetExprTestNode
		str          string
		size         uint32
		charsetNum   uint32
		plusNum      uint32
		minusNum     uint32
		branchNum    uint32
		hasWellKnown bool
	}{
		{[]charsetExprTestNode{}, "", 0, 0, 0, 0, 0, false},
		{[]charsetExprTestNode{{"alpha", true, []string{"a-z", "A-Z"}, CHARSET_OP_PLUS}}, "{alpha}", 52, 1, 1, 0, 1, true},
		{[]charsetExprTestNode{{"alpha", true, []string{"a-z", "A-Z"}, CHARSET_OP_MINUS}}, "", 0, 0, 0, 0, 0, false},
		{[]charsetExprTestNode{{"alpha", true, []string{"a-z", "A-Z"}, CHARSET_OP_PLUS}, {"digit", true, []string{"0-9"}, CHARSET_OP_PLUS}}, "{alpha} + {digit}", 62, 2, 2, 0, 2, true},
		{[]charsetExprTestNode{{"alpha", true, []string{"a-z", "A-Z"}, CHARSET_OP_PLUS}, {"digit", true, []string{"0-9"}, CHARSET_OP_MINUS}}, "{alpha} - {digit}", 52, 2, 1, 1, 2, true},
		{[]charsetExprTestNode{{"alpha", true, []string{"a-z", "A-Z"}, CHARSET_OP_PLUS}, {"token", false, []string{"a-b"}, CHARSET_OP_MINUS}}, "{alpha} - {a, b}", 50, 2, 1, 1, 3, true},
		{[]charsetExprTestNode{{"alpha", true, []string{"a-z", "A-Z"}, CHARSET_OP_MINUS}, {"token", false, []string{"a-b"}, CHARSET_OP_PLUS}}, "{a, b}", 2, 1, 1, 0, 2, false},
	}
	prefix := trace.CallerName(0)

	for i, v := range testdata {
		expr := buildCharsetExpr(v.expr)

		str := expr.String()
		if str != v.str {
			t.Errorf("%s[%d] failed: str = %s, wanted = %s\n", prefix, i, str, v.str)
		}

		if expr.CharsetSize() != v.size {
			t.Errorf("%s[%d] failed: size = %d, wanted = %d\n", prefix, i, expr.CharsetSize(), v.size)
		}

		if expr.CharsetNum() != v.charsetNum {
			t.Errorf("%s[%d] failed: charsetNum = %d, wanted = %d\n", prefix, i, expr.CharsetNum(), v.charsetNum)
		}

		if !expr.Empty() && v.charsetNum == 0 {
			t.Errorf("%s[%d] failed: should be empty\n", prefix, i)
		}

		if expr.Empty() && v.charsetNum > 0 {
			t.Errorf("%s[%d] failed: should not be empty\n", prefix, i)
		}

		if expr.PlusNum() != v.plusNum {
			t.Errorf("%s[%d] failed: plusNum = %d, wanted = %d\n", prefix, i, expr.PlusNum(), v.plusNum)
		}

		if expr.MinusNum() != v.minusNum {
			t.Errorf("%s[%d] failed: minusNum = %d, wanted = %d\n", prefix, i, expr.MinusNum(), v.minusNum)
		}

		if expr.BranchNum() != v.branchNum {
			t.Errorf("%s[%d] failed: branchNum = %d, wanted = %d\n", prefix, i, expr.BranchNum(), v.branchNum)
		}

		if !expr.HasWellKnownCharset() && v.hasWellKnown {
			t.Errorf("%s[%d] failed: should have wellknown charset\n", prefix, i)
		}

		if expr.HasWellKnownCharset() && !v.hasWellKnown {
			t.Errorf("%s[%d] failed: should not have wellknown charset\n", prefix, i)
		}
	}
}
