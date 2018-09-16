package peg

import (
	//"bytes"
	"fmt"
	"io"

	//"github.com/lioneagle/abnf/src/gen/peg_gen"

	"github.com/lioneagle/goutil/src/buffer"
)

type PegChoice struct {
	PegBase
	exprs []Peg
}

func (this *PegChoice) Type() uint32 { return PEG_CHOICE }

func NewPegChoice(name string) *PegChoice {
	expr := &PegChoice{}
	expr.SetName(name)
	return expr
}

func (this *PegChoice) Append(expr Peg) {
	this.exprs = append(this.exprs, expr)
}

func (this *PegChoice) PrintAsAbnf(w io.Writer) io.Writer {
	num := len(this.exprs)
	if num <= 0 {
		return w
	}

	if this.HasName() {
		fmt.Fprint(w, this.name)
		return w
	}

	if num == 1 {
		this.exprs[0].PrintAsAbnf(w)
		return w
	}

	fmt.Fprint(w, "(")
	this.exprs[0].PrintAsAbnf(w)
	for i := 1; i < num; i++ {
		fmt.Fprint(w, " / ")
		this.exprs[i].PrintAsAbnf(w)
	}
	fmt.Fprint(w, ")")

	return w
}

func (this *PegChoice) String() string {
	buf := buffer.NewByteBuffer(nil)
	this.PrintAsAbnf(buf)
	return buf.String()
}
