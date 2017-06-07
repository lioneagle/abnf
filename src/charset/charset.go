package charset

import (
	"basic"
	"container/list"
	//"fmt"
	//"os"
	//"errors"
	//"strconv"
)

/* Charset is a list of Range to represent a charset
 */
type Charset struct {
	list.List
	IsWellKnown bool
	Name        string
	size        int32
}

func (c *Charset) Empty() bool { return c.Len() == 0 }
func (c *Charset) Size() int32 { return c.size }

func (c *Charset) remove(e *list.Element) *list.Element {
	next := e.Next()
	c.Remove(e)
	return next
}

func (c *Charset) RemoveAll() {
	for e := c.Front(); e != nil; {
		e = c.remove(e)
	}
	c.size = 0
}

func (c *Charset) Contains(ch int32) bool {
	for e := c.Front(); e != nil; e = e.Next() {
		r := e.Value.(*Range)
		if r.Contains(ch) {
			return true
		}
	}
	return false
}

func (c *Charset) UniteCharset(c2 *Charset) {
	/*  @@TODO: 目前的算法是最简单的，但性能上要稍微差一些，以后再改进

		RangeList本身是有序且最大程度合并的，即不会存在[1,2)->[2,3)、[1,4)->[2,3)、[3,4)->[1,2)等情况，
	    因此在合并两个RangeList的时候，只需要依次检查两个RangeList的当前元素是否可以合并，但必须每个都
	    检查，因为可能存在这样的情况

	    range_list1: [1,2)->[3,4)->[5,6)->...
	    range_list1: [2,3)->[4,5)->[6,7)->...
	*/

	for e := c2.Front(); e != nil; e = e.Next() {
		r := e.Value.(*Range)
		c.UniteRange(r)
	}
}

func (c *Charset) UniteRange(r *Range) {

	r.Assert()

	if c.Empty() {
		c.PushBack(r)
		c.size += r.Size()
		return
	}

	var val *Range

	iter := c.Front()

	for {
		val = iter.Value.(*Range)

		if val.Low > r.High {
			c.InsertBefore(r, iter)
			c.size += r.Size()
			return
		}

		if val.High < r.Low {
			iter = iter.Next()
			if iter == nil {
				c.PushBack(r)
				c.size += r.Size()
				return
			}
		} else {
			break
		}
	}

	if val.Low >= r.Low {
		c.size += val.Low - r.Low
		val.Low = r.Low
	}

	if val.High >= r.High {
		return
	}

	c.size += r.High - val.High
	val.High = r.High

	newIter := iter.Next()

	for newIter != nil {
		newVal := newIter.Value.(*Range)
		if val.High < newVal.Low {
			break
		}

		if val.High <= newVal.High {
			c.size -= val.High - newVal.Low
			val.High = newVal.High

		} else {
			c.size -= newVal.High - newVal.Low
		}

		newIter = c.remove(newIter)
	}
}

func (c *Charset) UniteChar(ch int32) {
	c.UniteRange(&Range{ch, ch + 1})
}

func (c *Charset) Print(w basic.AbnfWriter) basic.AbnfWriter {

	if c.Empty() {
		return w
	}

	e := c.Front()
	val := e.Value.(*Range)
	val.PrintAsInt(w)
	e = e.Next()

	for ; e != nil; e = e.Next() {
		w.WriteString(", ")
		val = e.Value.(*Range)
		val.PrintAsInt(w)
	}
	return w
}

func (c *Charset) PrintAsChar(w basic.AbnfWriter) basic.AbnfWriter {

	if c.Empty() {
		return w
	}

	e := c.Front()
	val := e.Value.(*Range)
	val.PrintAsChar(w)
	e = e.Next()

	for ; e != nil; e = e.Next() {
		w.WriteString(", ")
		val = e.Value.(*Range)
		val.PrintAsChar(w)
	}
	return w
}

func (c *Charset) PrintEachChar(w basic.AbnfWriter) basic.AbnfWriter {

	if c.Empty() {
		return w
	}

	e := c.Front()
	val := e.Value.(*Range)
	val.PrintEachChar(w)
	e = e.Next()

	for ; e != nil; e = e.Next() {
		w.WriteString(", ")
		val = e.Value.(*Range)
		val.PrintEachChar(w)
	}
	return w
}

func (c *Charset) MakeFromBytes(str []byte) error {

	if len(str) == 0 {
		return nil
	}

	c.RemoveAll()

	var i int
	var low, high int32

	for i = 0; i < len(str); {

		low, i = basic.UnescapeChar(str, i)

		if i >= len(str) {
			c.UniteChar(low)
			return nil
		}

		if str[i] != '-' {
			c.UniteChar(low)
			continue
		}

		i++

		if i >= len(str) {
			c.UniteChar(low)
			c.UniteChar('-')
			return nil
		}

		high, i = basic.UnescapeChar(str, i)

		if low > high {
			low, high = high, low
		}

		c.UniteRange(&Range{low, high})

	}

	return nil
}
