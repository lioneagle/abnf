package peg

import (
	"io"

	"github.com/lioneagle/abnf/src/gen/peg_gen"
)

const (
	PEG_SEQUENCE uint32 = iota
	PEG_CHOICE
	PEG_CLOSURE
	PEG_AND_PREDICATE
	PEG_NOT_PREDICATE
	PEG_CHARSET
	PEG_STRING
	PEG_CONST_STRING
	PEG_KEYS
	PEG_IPV4
	PEG_IPV6
)

const (
	PEG_INFINITE_NUM = -1
)

type Peg interface {
	Type() uint32
	HasName() bool
	GetName() string
	SetName(name string)
	PrintAsAbnf(w io.Writer) io.Writer
	String() string
	GetTypeInfo() peg_gen.TypeInfo
	SetTypeInfo(info peg_gen.TypeInfo)
}

type PegBase struct {
	name     string
	typeinfo peg_gen.TypeInfo
}

func (this *PegBase) HasName() bool                     { return len(this.name) > 0 }
func (this *PegBase) GetName() string                   { return this.name }
func (this *PegBase) SetName(name string)               { this.name = name }
func (this *PegBase) GetTypeInfo() peg_gen.TypeInfo     { return this.typeinfo }
func (this *PegBase) SetTypeInfo(info peg_gen.TypeInfo) { this.typeinfo = info }
