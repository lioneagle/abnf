package charset

import (
	"basic"
	"bytes"
	"container/list"
	//"fmt"
	//"os"
	//"errors"
	//"strconv"
)

/* Charset is a list of Range to represent a charset
 */
type Charset struct {
	ranges      list.List
	IsWellKnown bool
	Name        string
	size        uint32
}

func NewCharset() *Charset         { return &Charset{} }
func (this *Charset) Empty() bool  { return this.ranges.Len() == 0 }
func (this *Charset) Size() uint32 { return this.size }

func (this *Charset) Equal(rhs *Charset) bool {
	if this.ranges.Len() != rhs.ranges.Len() {
		return false
	}

	iter1 := this.ranges.Front()
	iter2 := rhs.ranges.Front()

	for iter1 != nil {
		if !iter1.Value.(*Range).Equal(iter2.Value.(*Range)) {
			return false
		}
		iter1 = iter1.Next()
		iter2 = iter2.Next()
	}
	return true
}

func (this *Charset) remove(pos *list.Element) *list.Element {
	next := pos.Next()
	this.size -= pos.Value.(*Range).Size()
	this.ranges.Remove(pos)
	return next
}

func (this *Charset) appendRange(r *Range) {
	if r.Size() == 0 {
		return
	}
	this.ranges.PushBack(r.Clone())
	this.size += r.Size()
}

func (this *Charset) appendRanges(from, to *list.Element) {
	for iter := from; iter != to; iter = iter.Next() {
		this.appendRange(iter.Value.(*Range))
	}
}

func (this *Charset) RemoveAll() {
	for iter := this.ranges.Front(); iter != nil; {
		iter = this.remove(iter)
	}
}

func (this *Charset) Contains(ch int32) bool {
	for iter := this.ranges.Front(); iter != nil; iter = iter.Next() {
		r := iter.Value.(*Range)
		if r.Contains(ch) {
			return true
		}
	}
	return false
}

func (this *Charset) UniteCharset(rhs *Charset) {
	this.uniteRangeList(&rhs.ranges)
}

func (this *Charset) UniteRange(r *Range) {
	c := NewCharset()
	c.appendRange(r)
	this.UniteCharset(c)
}

func (this *Charset) UniteRangeSlice(ranges []Range) {
	for i := 0; i < len(ranges); i++ {
		this.UniteRange(&ranges[i])
	}
}

func (this *Charset) UniteChar(ch int32) {
	this.UniteRange(&Range{ch, ch + 1})
}

/* Range list of Charset is sorted, and each range is not intersected, so we just
   check each node of input range list and move iterator of this Charset without
   backtracking.

   The complexity is O(m+n), two range list are both traversed once.
*/
func (this *Charset) uniteRangeList(ranges *list.List) {
	iter1 := this.ranges.Front()
	iter2 := ranges.Front()

	for iter1 != nil && iter2 != nil {
		r2 := iter2.Value.(*Range)
		r1 := iter1.Value.(*Range)

		if r2.High < r1.Low {
			// insert between two nodes
			this.ranges.InsertBefore(r2.Clone(), iter1)
			this.size += r2.Size()
			iter2 = iter2.Next()
		} else if r2.Low > r1.High {
			iter1 = iter1.Next()
		} else {
			/* have intercetion */
			needMerge := false

			if r2.Low < r1.Low {
				this.size += uint32(r1.Low - r2.Low)
				r1.Low = r2.Low
				needMerge = true
			}
			if r1.High < r2.High {
				this.size += uint32(r2.High - r1.High)
				r1.High = r2.High
				needMerge = true
			}

			iter2 = iter2.Next()

			if needMerge {
				this.mergeFollowedRanges(iter1)
			}
		}
	}

	this.appendRanges(iter2, nil)
}

func (this *Charset) mergeFollowedRanges(pos *list.Element) {
	r1 := pos.Value.(*Range)
	for iter := pos.Next(); iter != nil; {
		r2 := iter.Value.(*Range)
		if r2.Low > r1.High {
			break
		}

		if r2.High >= r1.High {
			this.size += uint32(r2.High - r1.High)
			r1.High = r2.High
		}
		iter = this.remove(iter)
	}
}

func (this *Charset) DifferenceCharset(rhs *Charset) {
	this.differenceRangeList(&rhs.ranges)
}

func (this *Charset) DifferenceRange(r *Range) {
	c := NewCharset()
	c.appendRange(r)
	this.DifferenceCharset(c)
}

func (this *Charset) DifferenceChar(ch int32) {
	this.DifferenceRange(&Range{ch, ch + 1})
}

func (this *Charset) differenceRangeList(ranges *list.List) {
	iter1 := this.ranges.Front()
	iter2 := ranges.Front()

	for iter1 != nil && iter2 != nil {
		r1 := iter1.Value.(*Range)
		r2 := iter2.Value.(*Range)

		if r2.High <= r1.Low {
			iter2 = iter2.Next()
		} else if r1.High <= r2.Low {
			iter1 = iter1.Next()
		} else if r1.Low < r2.Low {
			if r1.High <= r2.High {
				this.size -= uint32(r1.High - r2.Low)
				r1.High = r2.Low
				iter1 = iter1.Next()
			} else {
				this.size -= r2.Size()
				this.ranges.InsertBefore(&Range{r1.Low, r2.Low}, iter1)
				r1.Low = r2.High
				iter2 = iter2.Next()
			}
		} else {
			if r1.High <= r2.High {
				iter1 = this.remove(iter1)
			} else {
				this.size -= uint32(r2.High - r1.Low)
				r1.Low = r2.High
				iter2 = iter2.Next()
			}
		}
	}
}

func (this *Charset) StringAsInt() string {
	buf := &bytes.Buffer{}
	this.PrintAsInt(buf)
	return buf.String()
}

func (this *Charset) StringAsChar() string {
	buf := &bytes.Buffer{}
	this.PrintAsChar(buf)
	return buf.String()
}

func (this *Charset) PrintAsInt(w basic.AbnfWriter) basic.AbnfWriter {
	return this.print(w, print_as_int)
}

func (this *Charset) PrintAsChar(w basic.AbnfWriter) basic.AbnfWriter {
	return this.print(w, print_as_char)
}

func (this *Charset) PrintEachChar(w basic.AbnfWriter) basic.AbnfWriter {
	return this.print(w, print_each_char)
}

const (
	print_as_int = iota
	print_as_char
	print_each_char
)

func (this *Charset) print(w basic.AbnfWriter, printType int) basic.AbnfWriter {
	for iter := this.ranges.Front(); iter != nil; iter = iter.Next() {
		if iter != this.ranges.Front() {
			w.WriteString(", ")
		}
		val := iter.Value.(*Range)
		switch printType {
		case print_as_int:
			val.PrintAsInt(w)
		case print_as_char:
			val.PrintAsChar(w)
		case print_each_char:
			val.PrintEachChar(w)
		}
	}
	return w
}

func (this *Charset) toString(printType int) string {
	buf := &bytes.Buffer{}
	this.print(buf, printType)
	return buf.String()
}

func (this *Charset) MakeFromBytes(str []byte) {
	this.RemoveAll()

	var i int
	var low, high int32

	for i = 0; i < len(str); {

		low, i = basic.UnescapeChar(str, i)

		if i >= len(str) {
			this.UniteChar(low)
			return
		}

		if str[i] != '-' {
			this.UniteChar(low)
			continue
		}

		i++

		if i >= len(str) {
			this.UniteChar(low)
			this.UniteChar('-')
			return
		}

		high, i = basic.UnescapeChar(str, i)

		if low > high {
			low, high = high, low
		}

		this.UniteRange(&Range{low, high})
	}
}

func (this *Charset) MakeFromBytesInverse(rangeAny *Range, str []byte) {
	this.UniteRange(rangeAny)
	c1 := NewCharset()
	c1.MakeFromBytes(str)
	this.DifferenceCharset(c1)
}
