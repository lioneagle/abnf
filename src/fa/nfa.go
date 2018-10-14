package fa

import (
	"fmt"
	"io"

	"github.com/lioneagle/abnf/src/charset"

	"github.com/lioneagle/goutil/src/buffer"
)

/* 由于采用Thompson构造法，因此每一个NFA的状态最多有两个迁移，且这两个迁移的字符不会同时为
 * 有效字符（即至少一个为epsilon），实际测试发现，Thompson构造法生成的DFA的状态比较多，影响
 * DFA生成和DFA状态最小化的效率。
 *
 * 主要是在生成多个字符（或者说一个字符集）各个元素或构成的NFA时会生成很多状态，
 * 因此考虑做一个改进，即将单个字符改为一个字符集，
 *
 *
 * 在构造NFA时，采用数组的方式存储NFA
 * 的状态（这是由于可以通过计算正则表达式的长度来计算NFA的最大状态数），用Thompson构造法，只
 * 会产生一个Final状态
 *
 */
type NfaTransition struct {
	isEpsiolon          bool
	isDefaultTransition bool
	charset             *charset.Charset
	charsetExpr         *charset.CharsetExpr
	actions             *ActionList
	destState           *NfaState
}

func NewNfaTransition() *NfaTransition {
	ret := &NfaTransition{}
	ret.charset = charset.NewCharset()
	ret.charsetExpr = charset.NewCharsetExpr()
	return ret
}

func (this *NfaTransition) SetDestSate(dest *NfaState) {
	this.destState = dest
}

func (this *NfaTransition) GetDestSate() *NfaState {
	return this.destState
}

func (this *NfaTransition) SetDefaultTransition() {
	this.isDefaultTransition = true
}

func (this *NfaTransition) SetNoneDefaultTransition() {
	this.isDefaultTransition = false
}

func (this *NfaTransition) SetNoneEpsilon() {
	this.isEpsiolon = false
}

func (this *NfaTransition) SetEpsilon() {
	this.isEpsiolon = true
}

func (this *NfaTransition) AddChar(ch byte) {
	this.charset.UniteChar(int32(ch))
}

func (this *NfaTransition) AddCharset(charset *charset.Charset) {
	this.charset.UniteCharset(charset)
}

func (this *NfaTransition) AddAction(action Action) {
	//TODO: need adding implement
}

func (this *NfaTransition) GetActions() *ActionList {
	return this.actions
}

func (this *NfaTransition) IsDefaultTransition() bool {
	return this.isDefaultTransition
}

func (this *NfaTransition) String() string {
	buf := buffer.NewByteBuffer(nil)
	this.Fprint(buf)
	return buf.String()
}

func (this *NfaTransition) Fprint(w io.Writer) {
	if this.isDefaultTransition {
		return
	}

	fmt.Fprint(w, "{ ")
	if this.isEpsiolon {
		fmt.Fprint(w, "epsilon")
	} else if !this.charsetExpr.Empty() {
		this.charsetExpr.Print(w)
	} else {
		this.charset.PrintEachChar(w)
	}

	if this.actions != nil && !this.actions.Empty() {
		fmt.Fprint(w, "<")
		this.actions.Fprint(w)
		fmt.Fprint(w, ">")
	}

	fmt.Println("this.destState =", this.destState)

	fmt.Fprintf(w, " } --> q%04d", this.destState.GetId())
}

type NfaState struct {
	id           uint32
	isFinal      bool
	entryActions ActionList
	exitActions  ActionList
}

func NewNfaState() *NfaState {
	return &NfaState{}
}

func (this *NfaState) GetId() uint32 {
	return this.id
}

func (this *NfaState) String() string {
	buf := buffer.NewByteBuffer(nil)
	this.Fprint(buf)
	return buf.String()
}

func (this *NfaState) Fprint(w io.Writer) {
	fmt.Fprintf(w, "q%04d", this.id)
	if this.isFinal {
		fmt.Fprintf(w, " (final)")
	}
}
