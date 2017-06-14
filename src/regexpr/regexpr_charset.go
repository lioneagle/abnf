package regexpr

import (
	"basic"
	"bytes"
	"charset"
	//"fmt"
)

type RegExprCharset struct {
	name    string
	charset charset.Charset
}

func NewRegExprCharset(name string, str []byte) *RegExprCharset {
	expr := &RegExprCharset{}
	expr.name = name
	expr.MakeFromBytes(str)
	return expr
}

func (this *RegExprCharset) Type() uint32        { return REG_EXPR_CHARSET }
func (this *RegExprCharset) HasName() bool       { return len(this.name) > 0 }
func (this *RegExprCharset) GetName() string     { return this.name }
func (this *RegExprCharset) SetName(name string) { this.name = name }

func (this *RegExprCharset) Print(w basic.AbnfWriter) basic.AbnfWriter {
	if this.charset.Empty() {
		return w
	}
	if this.HasName() {
		w.WriteString(this.name)
	} else if this.charset.Size() == 1 {
		w.WriteString("\"")
		this.charset.PrintAsChar(w)
		w.WriteString("\"")
	} else {
		w.WriteString("[")
		this.charset.PrintAsChar(w)
		w.WriteString("]")
	}
	return w
}

func (this *RegExprCharset) String() string {
	buf := &bytes.Buffer{}
	this.Print(buf)
	return buf.String()
}

func (this *RegExprCharset) MakeFromBytes(str []byte) {
	this.charset.MakeFromBytes(str)
}

func (this *RegExprCharset) MakeFromBytesInverse(rangeAny *charset.Range, str []byte) {
	this.charset.MakeFromBytesInverse(rangeAny, str)
}
