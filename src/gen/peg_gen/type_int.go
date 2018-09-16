package peg_gen

type TypeInfoInt struct {
	TypeInfoBase
	Size   int
	Signed bool
}

func NewTypeInfoInt(name string) *TypeInfoInt {
	info := &TypeInfoInt{}
	info.SetName(name)
	return info
}

func (this *TypeInfoInt) Kind() int               { return TYPE_KIND_INT }
func (this *TypeInfoInt) GetSize(align int) int   { return this.Size }
func (this *TypeInfoInt) CalcAlignSize() int      { return this.Size }
func (this *TypeInfoInt) GetAlignSize() int       { return this.Size }
func (this *TypeInfoInt) CalcPadNumber(align int) {}
func (this *TypeInfoInt) SetSize(size int)        { this.Size = size }
func (this *TypeInfoInt) SetSigned()              { this.Signed = true }
func (this *TypeInfoInt) SetUnsigned()            { this.Signed = false }
