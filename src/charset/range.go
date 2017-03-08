package charset

import (
	"fmt"
	//"os"
	//"errors"
	"strconv"
)

type ByteAndStringWriter interface {
	Write(p []byte) (n int, err error)
	WriteString(s string) (n int, err error)
}

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

func (r *Range) Print(w ByteAndStringWriter) ByteAndStringWriter {

	if r.Size() > 1 {
		w.WriteString(fmt.Sprintf("%d-%d", r.Low, r.High))
	} else {
		w.WriteString(fmt.Sprintf("%d", r.Low))
	}
	return w
}

func (r *Range) PrintAsChar(w ByteAndStringWriter) ByteAndStringWriter {

	PrintChar(w, r.Low)

	if r.Size() > 1 {
		w.WriteString("-")
		PrintChar(w, r.High)
	}
	return w
}

func (r *Range) PrintEachChar(w ByteAndStringWriter) ByteAndStringWriter {

	PrintChar(w, r.Low)
	for i := r.Low + 1; i < r.High; i++ {
		w.WriteString(", ")
		PrintChar(w, i)
	}
	return w
}

func PrintChar(w ByteAndStringWriter, ch int32) ByteAndStringWriter {
	switch ch {
	case '\'':
		w.WriteString("'")
	case '\n':
		w.WriteString("\\n")
	case '\t':
		w.WriteString("\\t")
	case '\v':
		w.WriteString("\\v")
	case '\b':
		w.WriteString("\\b")
	case '\r':
		w.WriteString("\\r")
	case '\f':
		w.WriteString("\\f")
	case '\a':
		w.WriteString("\\a")
	case '\\':
		w.WriteString("\\\\")
	default:
		if ch >= 0 && ch < 256 {
			if strconv.IsPrint(ch) && ch <= '~' {
				w.WriteString(fmt.Sprintf("%c", ch))
			} else {
				w.WriteString(fmt.Sprintf("\\x%02x", ch))
			}
		} else {
			w.WriteString(fmt.Sprintf("%d", ch))
		}

	}

	return w
}
