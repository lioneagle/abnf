package peg

import (
	//"bytes"
	"fmt"
	"io"

	//"github.com/lioneagle/abnf/src/gen/peg_gen"

	"github.com/lioneagle/goutil/src/buffer"
)

type PegClosure struct {
	PegBase
	min  int
	max  int
	expr Peg
}

func NewPegClosure(name string, expr Peg, minNum, maxNum int) *PegClosure {
	expr1 := &PegClosure{expr: expr, min: minNum, max: maxNum}
	expr1.SetName(name)
	return expr1
}

func (this *PegClosure) Type() uint32 { return PEG_CLOSURE }

func (this *PegClosure) PrintAsAbnf(w io.Writer) io.Writer {
	if this.HasName() {
		fmt.Fprint(w, this.name)
		return w
	}

	if this.min == 0 && this.max == 1 {
		fmt.Fprint(w, "[")
		this.expr.PrintAsAbnf(w)
		fmt.Fprint(w, "]")
		return w
	}

	if this.min > 0 && this.min != PEG_INFINITE_NUM {
		fmt.Fprint(w, this.min)
	}

	fmt.Fprint(w, "*")
	if this.max > 0 && this.max != PEG_INFINITE_NUM {
		fmt.Fprint(w, this.max)
	}

	if this.expr.HasName() {
		this.expr.PrintAsAbnf(w)
	} else {
		fmt.Fprint(w, "(")
		this.expr.PrintAsAbnf(w)
		fmt.Fprint(w, ")")
	}

	return w
}

func (this *PegClosure) String() string {
	buf := buffer.NewByteBuffer(nil)
	this.PrintAsAbnf(buf)
	return buf.String()
}
