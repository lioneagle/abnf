package peg_gen

import (
	//"fmt"
	//"io"

	_ "github.com/lioneagle/abnf/src/basic"
)

type TypeInfoChar struct {
	TypeInfoBase
}

func NewTypeInfoChar(name string) *TypeInfoChar {
	info := &TypeInfoChar{}
	info.SetName(name)
	return info
}

func (this *TypeInfoChar) Kind() int               { return TYPE_KIND_CHAR }
func (this *TypeInfoChar) GetSize(align int) int   { return 1 }
func (this *TypeInfoChar) CalcAlignSize() int      { return 1 }
func (this *TypeInfoChar) GetAlignSize() int       { return 1 }
func (this *TypeInfoChar) CalcPadNumber(align int) {}
