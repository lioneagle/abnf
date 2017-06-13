package regexpr

import (
	"basic"
	"bytes"
)

type RegExprCat struct {
	Name  string
	expr1 RegExpr
	expr2 RegExpr
}

func (this *RegExprCat) Type() uint32    { return REG_EXPR_CAT }
func (this *RegExprCat) HasName() bool   { return len(this.Name) == 0 }
func (this *RegExprCat) GetName() string { return this.Name }

func (this *RegExprCat) Print(w basic.AbnfWriter) basic.AbnfWriter {
	if this.HasName() {
		w.WriteString(this.Name)
	} else {
		this.expr1.Print(w)
		this.expr2.Print(w)
	}
	return w
}

func (this *RegExprCat) String() string {
	buf := &bytes.Buffer{}
	this.Print(buf)
	return buf.String()
}
