package peg

import (
	//"bytes"
	"fmt"
	"io"

	//"github.com/lioneagle/abnf/src/gen/peg_gen"

	"github.com/lioneagle/goutil/src/buffer"
)

type PegKeys struct {
	PegBase
	keys []*PegConstString
}

func NewPegKeys(name string) *PegKeys {
	expr := &PegKeys{}
	expr.SetName(name)
	return expr
}

func (this *PegKeys) Type() uint32 { return PEG_KEYS }

func (this *PegKeys) AppendKey(key *PegConstString) {
	this.keys = append(this.keys, key)
}

func (this *PegKeys) PrintAsAbnf(w io.Writer) io.Writer {
	if len(this.keys) <= 0 {
		return w
	}

	if this.HasName() {
		fmt.Fprint(w, this.name)
		return w
	}

	if len(this.keys) == 1 {
		return this.keys[0].PrintAsAbnf(w)
	}

	fmt.Fprint(w, "(")
	this.keys[0].PrintAsAbnf(w)

	for i := 1; i < len(this.keys); i++ {
		fmt.Fprint(w, " / ")
		this.keys[i].PrintAsAbnf(w)
	}

	fmt.Fprint(w, ")")

	return w
}

func (this *PegKeys) String() string {
	buf := buffer.NewByteBuffer(nil)
	this.PrintAsAbnf(buf)
	return buf.String()
}
