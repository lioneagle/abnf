package c

import (
	"fmt"
	"io"
	//"os"
	//"path/filepath"
	//"strings"

	"github.com/lioneagle/abnf/src/backends"
	"github.com/lioneagle/abnf/src/gen"
	"github.com/lioneagle/abnf/src/gen/peg_gen"
	//"github.com/lioneagle/abnf/src/gen/key_gen/key_cmp_gen"
	//"github.com/lioneagle/abnf/src/keys"

	//"github.com/lioneagle/goutil/src/chars"
	//"github.com/lioneagle/goutil/src/logger"
	//"github.com/lioneagle/goutil/src/times"
)

type TypeGeneratorForC struct {
	backends.CGeneratorBase

	config *peg_gen.Config
}

func NewTypeGeneratorForC(config *peg_gen.Config) *TypeGeneratorForC {
	ret := &TypeGeneratorForC{}
	ret.Indent.Init(0, 4)
	ret.Config = gen.NewCConfigBase()
	ret.Config.Indent = config.Indent
	ret.config = config
	return ret
}

func (this *TypeGeneratorForC) GetTypeCharName(typeinfo *peg_gen.TypeInfoChar) string {
	if typeinfo.HasName() {
		return typeinfo.GetName()
	}
	return "char"
}

func (this *TypeGeneratorForC) GetTypeIntName(typeinfo *peg_gen.TypeInfoInt) string {
	if typeinfo.HasName() {
		return typeinfo.GetName()
	}

	var name string

	if !typeinfo.Signed {
		name = "unsigned "
	}

	switch typeinfo.Size {
	case 1:
		name += "char"
	case 2:
		name += "short"
	case 4:
		name += "int"
	default:
		name += "long"
	}

	return name
}

func (this *TypeGeneratorForC) GetTypeStringName(typeinfo *peg_gen.TypeInfoString) string {
	if typeinfo.HasName() {
		return typeinfo.GetName()
	}
	return "char"
}

func (this *TypeGeneratorForC) GetTypeStructName(typeinfo *peg_gen.TypeInfoStruct) string {
	if typeinfo.HasName() {
		return typeinfo.GetName()
	}
	return ""
}

func (this *TypeGeneratorForC) GetTypeName(typeinfo peg_gen.TypeInfo) string {
	switch typeinfo.Kind() {
	case peg_gen.TYPE_KIND_CHAR:
		return this.GetTypeCharName(typeinfo.(*peg_gen.TypeInfoChar))
	case peg_gen.TYPE_KIND_STRING:
		return this.GetTypeStringName(typeinfo.(*peg_gen.TypeInfoString))
	case peg_gen.TYPE_KIND_INT:
		return this.GetTypeIntName(typeinfo.(*peg_gen.TypeInfoInt))
	case peg_gen.TYPE_KIND_STRUCT:
		return this.GetTypeStructName(typeinfo.(*peg_gen.TypeInfoStruct))
	default:
		return "unknown"
	}
}

func (this *TypeGeneratorForC) GetVarName(field *peg_gen.Var) string {
	if field.GetTypeInfo().Kind() == peg_gen.TYPE_KIND_STRING {
		stringInfo, _ := field.GetTypeInfo().(*peg_gen.TypeInfoString)
		return fmt.Sprintf("%s[%d]", field.GetName(), stringInfo.StringSize())
	}

	return fmt.Sprintf("%s", field.GetName())
}

func (this *TypeGeneratorForC) GenerateStruct(fields *peg_gen.TypeInfoStruct, w io.Writer) {
	if !fields.HasName() {
		this.Fprintfln(w, "struct unknown {};")
		return
	}

	padIndex := 1

	s := gen.NewStruct()
	s.Name = fields.GetName()

	for _, v := range fields.Fields {
		padNum := v.GetPadNumber()

		var field *gen.Var

		if padNum > 0 {
			field = gen.NewVar()
			field.TypeName = this.config.PadTypeName
			field.Name = fmt.Sprintf("pad%d[%d]", padIndex, padNum)
			padIndex++
			s.AppendField(field)
		}

		field = gen.NewVar()
		field.TypeName = this.GetTypeName(v.GetTypeInfo())
		field.Name = this.GetVarName(v)
		field.Comment = v.GetComment()
		s.AppendField(field)

	}

	if fields.Pad > 0 {
		field := gen.NewVar()
		field.TypeName = this.config.PadTypeName
		field.Name = fmt.Sprintf("pad%d[%d]", padIndex, fields.Pad)
		padIndex++
		s.AppendField(field)

	}

	this.GenerateStructDefine(w, s)
}
