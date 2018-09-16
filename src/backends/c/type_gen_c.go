package c

import (
	"fmt"
	"io"
	//"os"
	//"path/filepath"
	//"strings"

	"github.com/lioneagle/abnf/src/basic"
	"github.com/lioneagle/abnf/src/gen/peg_gen"
	//"github.com/lioneagle/abnf/src/gen/key_gen/key_cmp_gen"
	//"github.com/lioneagle/abnf/src/keys"

	"github.com/lioneagle/goutil/src/chars"
	//"github.com/lioneagle/goutil/src/logger"
	//"github.com/lioneagle/goutil/src/times"
)

type TypeGeneratorForC struct {
	chars.Indent
}

func NewTypeGeneratorForC() *TypeGeneratorForC {
	ret := &TypeGeneratorForC{}
	ret.Indent.Init(0, 4)
	return ret
}

func (this *TypeGeneratorForC) GetTypeCharName(config *peg_gen.Config, typeinfo *peg_gen.TypeInfoChar) string {
	if typeinfo.HasName() {
		return typeinfo.GetName()
	}
	return "char"
}

func (this *TypeGeneratorForC) GetTypeIntName(config *peg_gen.Config, typeinfo *peg_gen.TypeInfoInt) string {
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

func (this *TypeGeneratorForC) GetTypeStringName(config *peg_gen.Config, typeinfo *peg_gen.TypeInfoString) string {
	if typeinfo.HasName() {
		return typeinfo.GetName()
	}
	return "char"
}

func (this *TypeGeneratorForC) GetTypeStructName(config *peg_gen.Config, typeinfo *peg_gen.TypeInfoStruct) string {
	if typeinfo.HasName() {
		return typeinfo.GetName()
	}
	return ""
}

func (this *TypeGeneratorForC) GetTypeName(config *peg_gen.Config, typeinfo peg_gen.TypeInfo) string {
	switch typeinfo.Kind() {
	case peg_gen.TYPE_KIND_CHAR:
		return this.GetTypeCharName(config, typeinfo.(*peg_gen.TypeInfoChar))
	case peg_gen.TYPE_KIND_STRING:
		return this.GetTypeStringName(config, typeinfo.(*peg_gen.TypeInfoString))
	case peg_gen.TYPE_KIND_INT:
		return this.GetTypeIntName(config, typeinfo.(*peg_gen.TypeInfoInt))
	case peg_gen.TYPE_KIND_STRUCT:
		return this.GetTypeStructName(config, typeinfo.(*peg_gen.TypeInfoStruct))
	default:
		return "unknown"
	}
}

func (this *TypeGeneratorForC) GetVarName(config *peg_gen.Config, field *peg_gen.Var) string {
	if field.GetTypeInfo().Kind() == peg_gen.TYPE_KIND_STRING {
		stringInfo, _ := field.GetTypeInfo().(*peg_gen.TypeInfoString)
		return fmt.Sprintf("%s[%d]", field.GetName(), stringInfo.StringSize())
	}

	return fmt.Sprintf("%s", field.GetName())
}

func (this *TypeGeneratorForC) GenerateVar(config *peg_gen.Config, field *peg_gen.Var, interval int, w io.Writer) {
	this.Fprintf(w, "%s", this.GetTypeName(config, field.GetTypeInfo()))
	basic.PrintIndent(w, interval)
	fmt.Fprintf(w, "%s", this.GetVarName(config, field))
}

func (this *TypeGeneratorForC) GenerateStruct(config *peg_gen.Config, field *peg_gen.TypeInfoStruct, w io.Writer) {
	if !field.HasName() {
		this.Fprintfln(w, "struct unknown {};")
		return
	}

	this.Fprintf(w, "typedef struct tag_%s", field.GetName())
	this.generateLeftBrace(config, w, config.IndentOfBlock)

	max_type_name_len := this.getStructMaxTypeLen(config, field)
	max_member_name_len := this.getStructMaxMemberNameLen(config, field)

	padIndex := 1

	for _, v := range field.Member {

		padNum := v.GetPadNumber()

		if padNum > 0 {
			this.Fprintf(w, "%s", config.PadTypeName)
			basic.PrintIndent(w, max_type_name_len-len(config.PadTypeName)+2)
			fmt.Fprintf(w, "pad%d[%d];", padIndex, padNum)
			this.PrintReturn(w)
			padIndex++

		}
		this.GenerateVar(config, v, max_type_name_len-len(this.GetTypeName(config, v.GetTypeInfo()))+2, w)
		fmt.Fprintf(w, ";")
		if v.HasComment() {
			basic.PrintIndent(w, max_member_name_len-len(this.GetVarName(config, v))+2)
			fmt.Fprintf(w, "// %s", v.GetComment())
		}
		this.PrintReturn(w)
	}

	if field.Pad > 0 {
		this.Fprintf(w, "%s", config.PadTypeName)
		basic.PrintIndent(w, max_type_name_len-len(config.PadTypeName)+2)
		fmt.Fprintf(w, "pad%d[%d];", padIndex, field.Pad)
		this.PrintReturn(w)

	}

	this.Exit()
	this.Fprintfln(w, "}%s;", field.GetName())
}

func (this *TypeGeneratorForC) getStructMaxTypeLen(config *peg_gen.Config, typeinfo *peg_gen.TypeInfoStruct) int {
	max := 0

	for _, v := range typeinfo.Member {
		type_name_len := len(this.GetTypeName(config, v.GetTypeInfo()))
		if type_name_len > max {
			max = type_name_len
		}
	}

	if len(config.PadTypeName) > max {
		max = len(config.PadTypeName)
	}

	return max
}

func (this *TypeGeneratorForC) getStructMaxMemberNameLen(config *peg_gen.Config, typeinfo *peg_gen.TypeInfoStruct) int {
	max := 0
	pad_len := 0
	pad_index := 1

	for _, v := range typeinfo.Member {
		name_len := len(this.GetVarName(config, v))
		if name_len > max {
			max = name_len
		}
		if v.GetPadNumber() > 0 {
			pad_len = len(fmt.Sprintf("pad%d[%d]", pad_index, v.GetPadNumber()))
			if pad_len > max {
				max = pad_len
			}
			pad_index++
		}
	}

	if typeinfo.Pad > 0 {
		pad_len = len(fmt.Sprintf("pad%d[%d]", pad_index, typeinfo.Pad))
		if pad_len > max {
			max = pad_len
		}
	}

	return max
}

func (this *TypeGeneratorForC) generateLeftBrace(config *peg_gen.Config, w io.Writer, indent int) {
	if config.BraceAtNextLine {
		fmt.Fprintln(w)
		this.Fprintln(w, "{")
	} else {
		fmt.Fprintln(w, " {")
	}
	this.EnterIndent(indent)
}
