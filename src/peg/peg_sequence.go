package peg

import (
	//"bytes"
	"fmt"
	"io"

	//"github.com/lioneagle/abnf/src/gen/peg_gen"

	"github.com/lioneagle/goutil/src/buffer"
)

type PegSequence struct {
	PegBase
	exprs []Peg
}

func (this *PegSequence) Type() uint32 { return PEG_SEQUENCE }

func NewPegSequence(name string) *PegSequence {
	expr := &PegSequence{}
	expr.SetName(name)
	return expr
}

func (this *PegSequence) Append(expr Peg) {
	this.exprs = append(this.exprs, expr)
}

func (this *PegSequence) PrintAsAbnf(w io.Writer) io.Writer {
	num := len(this.exprs)

	if num <= 0 {
		return w
	}

	if this.HasName() {
		fmt.Fprint(w, this.name)
		return w
	}

	this.exprs[0].PrintAsAbnf(w)
	for i := 1; i < num; i++ {
		fmt.Fprint(w, " ")
		this.exprs[i].PrintAsAbnf(w)
	}
	return w
}

func (this *PegSequence) String() string {
	buf := buffer.NewByteBuffer(nil)
	this.PrintAsAbnf(buf)
	return buf.String()
}
