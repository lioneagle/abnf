package peg

import (
	//"bytes"
	"fmt"
	"io"

	"github.com/lioneagle/abnf/src/charset"
	//"github.com/lioneagle/abnf/src/gen/peg_gen"

	"github.com/lioneagle/goutil/src/buffer"
)

type PegCharset struct {
	PegBase
	charset charset.Charset
}

func NewPegCharset(name string, chars []byte) *PegCharset {
	expr := &PegCharset{}
	expr.SetName(name)
	expr.MakeFromBytes(chars)
	return expr
}

func (this *PegCharset) Type() uint32 { return PEG_CHARSET }

func (this *PegCharset) PrintAsAbnf(w io.Writer) io.Writer {
	if this.charset.Empty() {
		return w
	}
	if this.HasName() {
		fmt.Fprint(w, this.GetName())
	} else if this.charset.Size() == 1 {
		this.charset.PrintAsChar(w)
	} else {
		fmt.Fprint(w, "[")
		this.charset.PrintAsChar(w)
		fmt.Fprint(w, "]")
	}
	return w
}

func (this *PegCharset) String() string {
	buf := buffer.NewByteBuffer(nil)
	this.PrintAsAbnf(buf)
	return buf.String()
}

func (this *PegCharset) MakeFromBytes(str []byte) {
	this.charset.MakeFromBytes(str)
}

func (this *PegCharset) MakeFromBytesInverse(rangeAny *charset.Range, str []byte) {
	this.charset.MakeFromBytesInverse(rangeAny, str)
}
