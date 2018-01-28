package regexpr

import (
	"io"
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
	Print(w io.Writer) io.Writer
	String() string
}
