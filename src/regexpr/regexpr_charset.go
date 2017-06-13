package regexpr

import (
	"basic"
	"bytes"
	"charset"
	//"fmt"
)

type RegExprCharset struct {
	Name    string
	charset charset.Charset
}

func NewRegExprCharset() *RegExprCharset {
	return &RegExprCharset{}
}

func (this *RegExprCharset) Type() uint32    { return REG_EXPR_CHARSET }
func (this *RegExprCharset) HasName() bool   { return len(this.Name) > 0 }
func (this *RegExprCharset) GetName() string { return this.Name }

func (this *RegExprCharset) Print(w basic.AbnfWriter) basic.AbnfWriter {
	if this.HasName() {
		w.WriteString(this.Name)
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
