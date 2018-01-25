package regexpr

import (
	"github.com/lioneagle/abnf/src/basic"
)

const (
	REG_EXPR_CAT uint32 = iota
	REG_EXPR_OR
	REG_EXPR_ALT
	REG_EXPR_CLOSURE
	REG_EXPR_CHARSET
	REG_EXPR_RULE
)

type RegExpr interface {
	Type() uint32
	HasName() bool
	GetName() string
	SetName(name string)
	Print(w basic.AbnfWriter) basic.AbnfWriter
	String() string
}
