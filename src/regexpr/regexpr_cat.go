package regexpr

import (
	"bytes"
	"fmt"
	"io"
)

type RegExprCat struct {
	name  string
	expr1 RegExpr
	expr2 RegExpr
}

func (this *RegExprCat) Type() uint32        { return REG_EXPR_CAT }
func (this *RegExprCat) HasName() bool       { return len(this.name) > 0 }
func (this *RegExprCat) GetName() string     { return this.name }
func (this *RegExprCat) SetName(name string) { this.name = name }

func NewRegExprCat(name string, expr1, expr2 RegExpr) *RegExprCat {
	return &RegExprCat{name: name, expr1: expr1, expr2: expr2}
}

func (this *RegExprCat) Print(w io.Writer) io.Writer {
	if this.HasName() {
		fmt.Fprint(w, this.name)
	} else {
		this.expr1.Print(w)
		fmt.Fprint(w, " ")
		this.expr2.Print(w)
	}
	return w
}

func (this *RegExprCat) String() string {
	buf := &bytes.Buffer{}
	this.Print(buf)
	return buf.String()
}
