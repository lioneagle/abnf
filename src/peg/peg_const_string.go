package peg

import (
	//"bytes"
	"fmt"
	"io"

	"github.com/lioneagle/abnf/src/charset"
	//"github.com/lioneagle/abnf/src/gen/peg_gen"

	"github.com/lioneagle/goutil/src/buffer"
)

type PegConstString struct {
	PegBase
	value           []byte
	isCaseSensitive bool
	charset         charset.Charset
}

func NewPegConstString(name string, value []byte, isCaseSensitive bool) *PegConstString {
	expr := &PegConstString{value: value, isCaseSensitive: isCaseSensitive}
	expr.SetName(name)
	return expr
}

func (this *PegConstString) Type() uint32 { return PEG_CONST_STRING }

func (this *PegConstString) PrintAsAbnf(w io.Writer) io.Writer {
	if len(this.value) <= 0 {
		return w
	}

	if this.HasName() {
		fmt.Fprint(w, this.name)
		return w
	}

	if this.isCaseSensitive {
		fmt.Fprint(w, "'")
		fmt.Fprint(w, string(this.value))
		fmt.Fprint(w, "'")
	} else {
		fmt.Fprint(w, "\"")
		fmt.Fprint(w, string(this.value))
		fmt.Fprint(w, "\"")
	}
	return w
}

func (this *PegConstString) String() string {
	buf := buffer.NewByteBuffer(nil)
	this.PrintAsAbnf(buf)
	return buf.String()
}
