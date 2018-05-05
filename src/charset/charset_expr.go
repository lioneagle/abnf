package charset

import (
	"bytes"
	"container/list"
	"fmt"
	"io"
)

const (
	CHARSET_OP_PLUS uint32 = iota
	CHARSET_OP_MINUS
)

type CharsetExprNode struct {
	charset *Charset
	op      uint32
}

type CharsetExpr struct {
	Name     string
	charsets list.List
	charset  Charset
	plusNum  uint32
	minusNum uint32
}

func NewCharsetExpr() *CharsetExpr {
	return &CharsetExpr{}
}

func (this *CharsetExpr) Empty() bool                        { return this.charsets.Len() == 0 }
func (this *CharsetExpr) PlusNum() uint32                    { return this.plusNum }
func (this *CharsetExpr) MinusNum() uint32                   { return this.minusNum }
func (this *CharsetExpr) CharsetNum() uint32                 { return uint32(this.charsets.Len()) }
func (this *CharsetExpr) CharsetSize() uint32                { return this.charset.Size() }
func (this *CharsetExpr) CharsetEqual(rhs *CharsetExpr) bool { return this.charset.Equal(&rhs.charset) }

func (this *CharsetExpr) HasWellKnownCharset() bool {
	for iter := this.charsets.Front(); iter != nil; iter = iter.Next() {
		if iter.Value.(*CharsetExprNode).charset.IsWellKnown {
			return true
		}
	}
	return false
}

func (this *CharsetExpr) BranchNum() (num uint32) {
	for iter := this.charsets.Front(); iter != nil; iter = iter.Next() {
		if iter.Value.(*CharsetExprNode).charset.IsWellKnown {
			num++
		} else {
			num += iter.Value.(*CharsetExprNode).charset.Size()
		}
	}
	return num
}

func (this *CharsetExpr) Plus(charset *Charset) {
	node := &CharsetExprNode{charset: charset, op: CHARSET_OP_PLUS}
	this.charsets.PushBack(node)
	this.charset.UniteCharset(charset)
	this.plusNum++
}

func (this *CharsetExpr) Minus(charset *Charset) {
	if this.charsets.Len() == 0 {
		return
	}
	node := &CharsetExprNode{charset: charset, op: CHARSET_OP_MINUS}
	this.charsets.PushBack(node)
	this.charset.DifferenceCharset(charset)
	this.minusNum++
}

func (this *CharsetExpr) String() string {
	buf := &bytes.Buffer{}
	this.Print(buf)
	return buf.String()
}

func (this *CharsetExpr) Print(w io.Writer) io.Writer {
	for iter := this.charsets.Front(); iter != nil; iter = iter.Next() {
		val := iter.Value.(*CharsetExprNode)
		if iter == this.charsets.Front() {
			fmt.Fprint(w, "{")
		} else {
			if val.op == CHARSET_OP_PLUS {
				fmt.Fprint(w, " + {")
			} else {
				fmt.Fprint(w, " - {")
			}
		}

		if val.charset.IsWellKnown {
			fmt.Fprint(w, val.charset.Name)
		} else {
			val.charset.PrintEachChar(w)
		}

		fmt.Fprint(w, "}")
	}
	return w
}
