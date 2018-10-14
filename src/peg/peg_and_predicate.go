package peg

import (
	//"bytes"
	"fmt"
	"io"

	//"github.com/lioneagle/abnf/src/gen/peg_gen"

	"github.com/lioneagle/goutil/src/buffer"
)

type PegAndPredicate struct {
	PegBase
	predicate Peg
}

func NewPegAndPredicate(name string, predicate Peg) *PegAndPredicate {
	expr1 := &PegAndPredicate{predicate: predicate}
	expr1.SetName(name)
	return expr1
}

func (this *PegAndPredicate) Type() uint32 { return PEG_AND_PREDICATE }

func (this *PegAndPredicate) PrintAsAbnf(w io.Writer) io.Writer {
	if this.HasName() {
		fmt.Fprint(w, this.name)
		return w
	}

	fmt.Fprint(w, "&")

	if this.predicate.HasName() {
		this.predicate.PrintAsAbnf(w)
	} else {
		fmt.Fprint(w, "(")
		this.predicate.PrintAsAbnf(w)
		fmt.Fprint(w, ")")
	}

	return w
}

func (this *PegAndPredicate) String() string {
	buf := buffer.NewByteBuffer(nil)
	this.PrintAsAbnf(buf)
	return buf.String()
}
