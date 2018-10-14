package peg

import (
	//"bytes"
	"fmt"
	"io"

	"github.com/lioneagle/abnf/src/charset"
	//"github.com/lioneagle/abnf/src/gen/peg_gen"

	"github.com/lioneagle/goutil/src/buffer"
)

type PegString struct {
	PegBase
	min     int
	max     int
	charset charset.Charset
}

func NewPegString(name string, chars []byte, minNum, maxNum int) *PegString {
	expr := &PegString{min: minNum, max: maxNum}
	expr.SetName(name)
	expr.MakeFromBytes(chars)
	return expr
}

func (this *PegString) Type() uint32 { return PEG_STRING }

func (this *PegString) PrintAsAbnf(w io.Writer) io.Writer {
	if this.charset.Empty() {
		return w
	}

	if this.HasName() {
		fmt.Fprint(w, this.name)
		return w
	}

	if this.min > 0 && this.min != PEG_INFINITE_NUM {
		fmt.Fprint(w, this.min)
	}

	fmt.Fprint(w, "*")
	if this.max > 0 && this.max != PEG_INFINITE_NUM {
		fmt.Fprint(w, this.max)
	}

	if this.charset.Size() == 1 {
		this.charset.PrintAsString(w)
	} else {
		fmt.Fprint(w, "[")
		this.charset.PrintAsString(w)
		fmt.Fprint(w, "]")
	}
	return w
}

func (this *PegString) String() string {
	buf := buffer.NewByteBuffer(nil)
	this.PrintAsAbnf(buf)
	return buf.String()
}

func (this *PegString) MakeFromBytes(str []byte) {
	this.charset.MakeFromBytes(str)
}

func (this *PegString) MakeFromBytesInverse(rangeAny *charset.Range, str []byte) {
	this.charset.MakeFromBytesInverse(rangeAny, str)
}
