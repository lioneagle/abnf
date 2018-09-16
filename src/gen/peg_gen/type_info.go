package peg_gen

import (
	_ "io"
)

const (
	TYPE_KIND_CHAR = iota
	TYPE_KIND_INT
	TYPE_KIND_STRING
	TYPE_KIND_STATIC_ARRAY
	TYPE_KIND_DYNAMIC_ARRAY
	TYPE_KIND_PTR
	TYPE_KIND_STRUCT
	TYPE_KIND_ENUM
	TYPE_KIND_IPV4
	TYPE_KIND_IPV6
	TYPE_KIND_IP
)

type TypeInfo interface {
	Kind() int
	GetSize(align int) int
	CalcAlignSize() int
	GetAlignSize() int
	HasName() bool
	GetName() string
	SetName(name string)
}

type TypeInfoBase struct {
	name string
}

func (this *TypeInfoBase) HasName() bool       { return len(this.name) > 0 }
func (this *TypeInfoBase) GetName() string     { return this.name }
func (this *TypeInfoBase) SetName(name string) { this.name = name }
