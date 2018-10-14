package regexpr

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lioneagle/abnf/src/charset"
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

func (this *RegExprCharset) Print(w io.Writer) io.Writer {
	if this.charset.Empty() {
		return w
	}
	if this.HasName() {
		fmt.Fprint(w, this.name)
	} else if this.charset.Size() == 1 {
		this.charset.PrintAsString(w)
	} else {
		fmt.Fprint(w, "[")
		this.charset.PrintAsString(w)
		fmt.Fprint(w, "]")
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
