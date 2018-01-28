package regexpr

import (
	"bytes"
	"fmt"
	"io"
)

type RegExprOr struct {
	name  string
	expr1 RegExpr
	expr2 RegExpr
}

func (this *RegExprOr) Type() uint32        { return REG_EXPR_CAT }
func (this *RegExprOr) HasName() bool       { return len(this.name) > 0 }
func (this *RegExprOr) GetName() string     { return this.name }
func (this *RegExprOr) SetName(name string) { this.name = name }

func NewRegExprOr(name string, expr1, expr2 RegExpr) *RegExprOr {
	return &RegExprOr{name: name, expr1: expr1, expr2: expr2}
}

func (this *RegExprOr) Print(w io.Writer) io.Writer {
	if this.HasName() {
		fmt.Fprint(w, this.name)
	} else {
		this.expr1.Print(w)
		fmt.Fprint(w, " | ")
		this.expr2.Print(w)
	}
	return w
}

func (this *RegExprOr) String() string {
	buf := &bytes.Buffer{}
	this.Print(buf)
	return buf.String()
}
