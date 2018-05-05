package fa

import (
	"io"

	"github.com/lioneagle/abnf/src/charset"
)

type State interface {
	IsMarked() bool
	SetMarked()
	SetUnmarked()

	SetFinal()
	SetUnfinal()
	IsFinal() bool

	HasEntryActions() bool
	HasExitActions() bool
	AddEntryAction(action Action)
	AddExitAction(action Action)

	GetEntryActions() ActionList
	GetExitActions() ActionList

	AddTransition(ch byte, destState State)
	GetDestState(ch byte) State
}

type Transition interface {
	SetDestSate(dest State)
	GetDestSate() State

	AddChar(ch byte)
	AddCharset(charset *charset.Charset)
	GetCharset() *charset.Charset
	GetCharsetExpr() *charset.CharsetExpr

	AddAction(action Action)
	GetActions() *ActionList
	IsDefaultTransition() bool

	String() string
	Fprint(w io.Writer)
}
