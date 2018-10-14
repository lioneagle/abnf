package peg_gen

import (
	//"fmt"

	"github.com/lioneagle/abnf/src/basic"
)

type Var struct {
	TypeInfo TypeInfo
	Name     string
	Comment  string
	Pad      int
}

func NewVar(typeInfo TypeInfo, name, comment string) *Var {
	return &Var{TypeInfo: typeInfo, Name: name, Comment: comment}
}

func (this *Var) GetTypeInfo() TypeInfo { return this.TypeInfo }
func (this *Var) GetName() string       { return this.Name }
func (this *Var) GetComment() string    { return this.Comment }
func (this *Var) HasComment() bool      { return len(this.Comment) > 0 }

func (this *Var) GetPadNumber() int    { return this.Pad }
func (this *Var) SetPadNumber(pad int) { this.Pad = pad }

type TypeInfoStruct struct {
	TypeInfoBase
	Align  int
	Size   int
	Pad    int
	Fields []*Var
}

func NewTypeInfoStruct(name string) *TypeInfoStruct {
	info := &TypeInfoStruct{}
	info.SetName(name)
	return info
}

func (this *TypeInfoStruct) Kind() int             { return TYPE_KIND_STRUCT }
func (this *TypeInfoStruct) GetAlignSize() int     { return this.Align }
func (this *TypeInfoStruct) GetSize(align int) int { return this.Size }

func (this *TypeInfoStruct) GetPadNumber() int    { return this.Pad }
func (this *TypeInfoStruct) SetPadNumber(pad int) { this.Pad = pad }

func (this *TypeInfoStruct) AppendMember(typeInfo TypeInfo, name, comment string) {
	this.Fields = append(this.Fields, NewVar(typeInfo, name, comment))
}

func (this *TypeInfoStruct) CalcAlignSize() int {
	max_align := 0
	for _, v := range this.Fields {
		align := v.GetTypeInfo().CalcAlignSize()
		//fmt.Printf("CalcAlignSize: name = %s, align = %d\n", v.Name(), align)
		if max_align < align {
			max_align = align
		}
	}
	this.Align = max_align
	//fmt.Printf("CalcAlignSize: name = %s, max_align = %d\n", this.GetName(), max_align)
	return max_align
}

func (this *TypeInfoStruct) CalcPadNumber(align int) int {
	self_align := this.GetAlignSize()
	if self_align > align {
		align = self_align
	}

	size := 0
	for _, v := range this.Fields {
		member_align := v.GetTypeInfo().GetAlignSize()
		//fmt.Printf("CalcPadNumber: name = %s, member_align = %d, size = %d\n", v.Name(), member_align, size)

		if size%member_align != 0 {
			aligned := basic.RoundToAlign(size, member_align)
			v.SetPadNumber(aligned - size)
			//fmt.Printf("CalcPadNumber: aligned = %d, size = %d, member_align = %d\n", aligned, size, member_align)
			size = aligned
		}

		//fmt.Printf("CalcPadNumber: name = %s, kind = %d\n", v.GetName(), v.GetTypeInfo().Kind())

		var size1 int
		if v.GetTypeInfo().Kind() != TYPE_KIND_STRUCT {
			size1 = v.GetTypeInfo().GetSize(align)
			//fmt.Printf("CalcPadNumber: name = %s, size = %d\n", v.GetName(), size1)
		} else {
			s, _ := v.GetTypeInfo().(*TypeInfoStruct)
			size1 = s.CalcPadNumber(align)
		}

		size += size1
	}

	//fmt.Printf("CalcPadNumber: name = %s, size = %d, align = %d\n", this.GetName(), size, align)

	if size%self_align != 0 {
		aligned := basic.RoundToAlign(size, align)
		this.SetPadNumber(aligned - size)
		size = aligned
	}

	this.Size = size

	return size
}
