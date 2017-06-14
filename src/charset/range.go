package charset

import (
	"basic"
	"bytes"
	"fmt"
	//"os"
	//"errors"
)

/* Range represent a range of [low,high)
 */
type Range struct {
	Low  int32
	High int32
}

func (this *Range) Contains(ch int32) bool   { return (this.Low <= ch) && (ch < this.High) }
func (this *Range) Size() uint32             { return uint32(this.High - this.Low) }
func (this *Range) Equal(r2 *Range) bool     { return this.Low == r2.Low && this.High == r2.High }
func (this *Range) Less(r2 *Range) bool      { return this.Low < r2.Low }
func (this *Range) LessEqual(r2 *Range) bool { return this.Low <= r2.Low }
func (this *Range) Assert() {
	if this.Low > this.High {
		panic(fmt.Sprintf("Range %v: Low > High", this))
	}
}

func (this *Range) Clone() *Range {
	return &Range{this.Low, this.High}
}

func (this *Range) StringAsInt() string {
	buf := &bytes.Buffer{}
	this.PrintAsInt(buf)
	return buf.String()
}

func (this *Range) PrintAsInt(w basic.AbnfWriter) basic.AbnfWriter {
	if this.Size() == 0 {
		return w
	}
	w.WriteString(fmt.Sprintf("%d", this.Low))
	if this.Size() > 1 {
		w.WriteString(fmt.Sprintf("-%d", this.High-1))
	}
	return w
}

func (this *Range) PrintAsChar(w basic.AbnfWriter) basic.AbnfWriter {
	if this.Size() == 0 {
		return w
	}
	basic.PrintIntAsChar(w, this.Low)
	if this.Size() > 1 {
		w.WriteString("-")
		basic.PrintIntAsChar(w, this.High-1)
	}
	return w
}

func (this *Range) PrintEachChar(w basic.AbnfWriter) basic.AbnfWriter {
	if this.Size() == 0 {
		return w
	}
	basic.PrintIntAsChar(w, this.Low)
	for i := this.Low + 1; i < this.High; i++ {
		w.WriteString(", ")
		basic.PrintIntAsChar(w, i)
	}
	return w
}
