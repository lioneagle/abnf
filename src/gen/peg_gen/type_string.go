package peg_gen

type TypeInfoString struct {
	TypeInfoBase
	Size int
}

func NewTypeInfoString(name string) *TypeInfoString {
	info := &TypeInfoString{}
	info.SetName(name)
	return info
}

func (this *TypeInfoString) Kind() int               { return TYPE_KIND_STRING }
func (this *TypeInfoString) GetSize(align int) int   { return this.Size }
func (this *TypeInfoString) CalcAlignSize() int      { return 1 }
func (this *TypeInfoString) GetAlignSize() int       { return 1 }
func (this *TypeInfoString) CalcPadNumber(align int) {}
func (this *TypeInfoString) SetSize(size int)        { this.Size = size }

func (this *TypeInfoString) StringSize() int { return this.Size }
