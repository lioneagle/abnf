package regexpr

import (
	"basic"
	"bytes"
)

type RegExprAlt struct {
	name string
	expr RegExpr
}

func (this *RegExprAlt) Type() uint32        { return REG_EXPR_ALT }
func (this *RegExprAlt) HasName() bool       { return len(this.name) > 0 }
func (this *RegExprAlt) GetName() string     { return this.name }
func (this *RegExprAlt) SetName(name string) { this.name = name }

func NewRegExprAlt(name string, expr RegExpr) *RegExprAlt {
	return &RegExprAlt{name: name, expr: expr}
}

func (this *RegExprAlt) Print(w basic.AbnfWriter) basic.AbnfWriter {
	if this.HasName() {
		w.WriteString(this.name)
	} else if this.expr.HasName() {
		this.expr.Print(w)
		w.WriteString("?")
	} else {
		w.WriteString("(")
		this.expr.Print(w)
		w.WriteString(")?")
	}
	return w
}

func (this *RegExprAlt) String() string {
	buf := &bytes.Buffer{}
	this.Print(buf)
	return buf.String()
}
