package charset

import (
	"basic"
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

func (r *Range) Contains(ch int32) bool   { return (r.Low <= ch) && (ch < r.High) }
func (r *Range) Size() int32              { return r.High - r.Low }
func (r *Range) Equal(r2 *Range) bool     { return r.Low == r2.Low && r.High == r2.High }
func (r *Range) Less(r2 *Range) bool      { return r.Low < r2.Low }
func (r *Range) LessEqual(r2 *Range) bool { return r.Low <= r2.Low }
func (r *Range) Assert() {
	if r.Low > r.High {
		panic(fmt.Sprintf("Range %v: Low > High", r))
	}
}

func (r *Range) PrintAsInt(w basic.AbnfWriter) basic.AbnfWriter {

	if r.Size() > 1 {
		w.WriteString(fmt.Sprintf("%d-%d", r.Low, r.High))
	} else if r.Size() == 1 {
		w.WriteString(fmt.Sprintf("%d", r.Low))
	}
	return w
}

func (r *Range) PrintAsChar(w basic.AbnfWriter) basic.AbnfWriter {

	basic.PrintIntAsChar(w, r.Low)

	if r.Size() > 1 {
		w.WriteString("-")
		basic.PrintIntAsChar(w, r.High)
	}
	return w
}

func (r *Range) PrintEachChar(w basic.AbnfWriter) basic.AbnfWriter {

	basic.PrintIntAsChar(w, r.Low)
	for i := r.Low + 1; i < r.High; i++ {
		w.WriteString(", ")
		basic.PrintIntAsChar(w, i)
	}
	return w
}
