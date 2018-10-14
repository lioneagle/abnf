package backends

import (
	"fmt"
	"io"
	//"os"
	"strings"

	"github.com/lioneagle/abnf/src/basic"
	"github.com/lioneagle/abnf/src/gen"

	"github.com/lioneagle/goutil/src/chars"
)

type GolangGeneratorBase struct {
	chars.Indent
	Config *gen.GolangConfigBase
}

func NewGolangGeneratorBase() *GolangGeneratorBase {
	gen := &GolangGeneratorBase{}
	gen.Indent.Init(0, 4)
	return gen
}

func (this *GolangGeneratorBase) GenerateFunctionDeclare(w io.Writer, f *gen.Function) {
}

func (this *GolangGeneratorBase) GenerateFunctionDefine(w io.Writer, f *gen.Function) {
	if len(f.Comment) > 0 {
		this.GenerateMultiLineComment(w, f.Comment)
	}

	this.Fprintf(w, "func %s(", f.Name)
	this.GenerateParamList(w, f.Params)
	fmt.Fprintf(w, ") %s", f.ReturnType)
}

func (this *GolangGeneratorBase) GenerateVarDeclare(w io.Writer, v *gen.Var, typeInterval, nameInterval int) {
	this.Fprintf(w, v.Name)
	basic.PrintIndent(w, nameInterval)
	fmt.Fprintf(w, "%s", v.TypeName)

	if len(v.Comment) > 0 {
		basic.PrintIndent(w, typeInterval+this.Config.Indent.Comment)
		this.GenerateSingleLineCommentWithoutIndent(w, v.Comment)
	} else {
		this.PrintReturn(w)
	}
}

func (this *GolangGeneratorBase) GenerateVarDefine(w io.Writer, v *gen.Var, typeInterval, nameInterval int) {
	this.Fprintf(w, "var %s", v.Name)
	basic.PrintIndent(w, nameInterval)
	fmt.Fprint(w, v.TypeName)
	if len(v.InitValue) > 0 {
		basic.PrintIndent(w, typeInterval+this.Config.Indent.Assign)
		fmt.Fprintf(w, "= %s", v.InitValue)
	}

	if len(v.Comment) > 0 {
		basic.PrintIndent(w, nameInterval+this.Config.Indent.Comment)
		this.GenerateSingleLineCommentWithoutIndent(w, v.Comment)
	} else {
		this.PrintReturn(w)
	}
}

func (this *GolangGeneratorBase) GenerateVarEnum(w io.Writer, v *gen.Var, nameInterval, valueInterval int) {
	this.Fprint(w, v.Name)
	if len(v.InitValue) > 0 {
		basic.PrintIndent(w, nameInterval+this.Config.Indent.Assign)
		fmt.Fprintf(w, "= %s", v.InitValue)
	}

	if len(v.Comment) > 0 {
		basic.PrintIndent(w, valueInterval+this.Config.Indent.Comment)
		this.GenerateSingleLineCommentWithoutIndent(w, v.Comment)
	}
}

func (this *GolangGeneratorBase) GenerateVarParam(w io.Writer, v *gen.Var) {
	fmt.Fprintf(w, "%s %s", v.Name, v.TypeName)
}

func (this *GolangGeneratorBase) GenerateParamList(w io.Writer, params *gen.ParamList) {
	if len(params.Params) <= 0 {
		return
	}

	this.GenerateVarParam(w, params.Params[0])

	if this.Config.ParamsInOneLine {
		for i := 1; i < len(params.Params); i++ {
			fmt.Fprintf(w, ", %s %s", params.Params[i].Name, params.Params[i].TypeName)
		}
	} else {
		this.EnterIndent(this.Config.Indent.FuncParam)
		for i := 1; i < len(params.Params); i++ {
			fmt.Fprint(w, ",")
			this.PrintReturn(w)
			this.Fprintf(w, "%s %s", params.Params[i].Name, params.Params[i].TypeName)
		}
		this.Exit()
	}
}

func (this *GolangGeneratorBase) GenerateMultiLineComment(w io.Writer, comment string) {
	lines := strings.Split(comment, "\n")
	if len(lines) <= 0 {
		return
	}

	newLines := make([]string, 0)
	max_len := 0

	for i := 0; i < len(lines); i++ {
		lines[i] = strings.TrimSpace(lines[i])
		if len(lines[i]) > 0 {
			newLines = append(newLines, lines[i])
		}

		if len(lines[i]) > max_len {
			max_len = len(lines[i])
		}
	}

	if len(newLines) <= 0 {
		return
	}

	if len(newLines) == 1 {
		this.Fprintfln(w, "/* %s */", newLines[0])
	} else if this.Config.MultiLineCommentDecorate {
		this.Fprint(w, "/*")
		basic.PrintChars(w, '*', max_len+2)
		this.PrintReturn(w)

		for i := 0; i < len(newLines); i++ {
			fmt.Fprintf(w, " * %s", newLines[i])
			this.PrintReturn(w)
		}
		fmt.Fprint(w, " ")
		basic.PrintChars(w, '*', max_len+3)
		this.Fprintln(w, "*/")
	} else {
		this.Fprint(w, "/* ")
		fmt.Fprint(w, newLines[0])
		this.PrintReturn(w)

		for i := 1; i < len(newLines); i++ {
			fmt.Fprintf(w, " * %s", newLines[i])
			this.PrintReturn(w)
		}
		this.Fprintln(w, " */")
	}
}

func (this *GolangGeneratorBase) GenerateSingleLineComment(w io.Writer, comment string) {
	if this.Config.VarUseSingleLineComment {
		this.Fprintfln(w, "// %s", comment)
	} else {
		this.Fprintfln(w, "/* %s */", comment)
	}
}

func (this *GolangGeneratorBase) GenerateSingleLineCommentWithoutIndent(w io.Writer, comment string) {
	if this.Config.VarUseSingleLineComment {
		fmt.Fprintf(w, "// %s", comment)
	} else {
		fmt.Fprintf(w, "/* %s */", comment)
	}
	this.PrintReturn(w)
}

func (this *GolangGeneratorBase) GenerateStructDefine(w io.Writer, s *gen.Struct) {
	this.Fprintf(w, "type %s struct", s.Name)
	this.GenerateBlockBegin(w, this.Config.Indent.Struct)

	maxTypeNameLen := s.Fields.GetMaxTypeNameLen()
	maxNameLen := s.Fields.GetMaxNameLen()

	for _, v := range s.Fields.Params {
		this.GenerateVarDeclare(w, v, maxTypeNameLen-len(v.TypeName), maxNameLen-len(v.Name)+1)
	}
	this.Exit()
	this.Fprintfln(w, "}")
}

func (this *GolangGeneratorBase) GenerateEnumDefine(w io.Writer, e *gen.Enum) {
	this.Fprintf(w, "const (")
	this.EnterIndent(this.Config.Indent.Enum)
	this.PrintReturn(w)

	maxNameLen := e.Enums.GetMaxNameLen()
	maxValueLen := e.Enums.GetMaxValueLen()

	for _, v := range e.Enums.Params {
		this.GenerateVarEnum(w, v, maxNameLen-len(v.Name), maxValueLen-len(v.InitValue))
	}
	this.Exit()
	this.Fprintfln(w, ")")
}

func (this *GolangGeneratorBase) GenerateSwitchBegin(w io.Writer, format string, args ...interface{}) {
	this.Fprint(w, "switch ")
	fmt.Fprintf(w, format, args...)

	this.GenerateBlockBegin(w, this.Config.Indent.Switch)
}

func (this *GolangGeneratorBase) GenerateSwitchEnd(w io.Writer) {
	this.GenerateBlockEnd(w)
}

func (this *GolangGeneratorBase) GenerateCaseBegin(w io.Writer, format string, args ...interface{}) {
	this.Fprint(w, "case ")
	fmt.Fprintf(w, format, args...)
	fmt.Fprint(w, ":")
	this.EnterIndent(this.Config.Indent.Case)
	this.PrintReturn(w)
}

func (this *GolangGeneratorBase) GenerateCaseEnd(w io.Writer) {
	this.Exit()
	this.PrintReturn(w)
}

func (this *GolangGeneratorBase) GenerateWhileBegin(w io.Writer, format string, args ...interface{}) {
}

func (this *GolangGeneratorBase) GenerateWhileEnd(w io.Writer) {
}

func (this *GolangGeneratorBase) GenerateDoWhileBegin(w io.Writer) {
}

func (this *GolangGeneratorBase) GenerateDoWhileEnd(w io.Writer, format string, args ...interface{}) {
}

func (this *GolangGeneratorBase) GenerateForBegin(w io.Writer, format string, args ...interface{}) {
	this.Fprint(w, "for ")
	fmt.Fprintf(w, format, args...)
	this.GenerateBlockBegin(w, this.Config.Indent.For)
}

func (this *GolangGeneratorBase) GenerateForEnd(w io.Writer) {
	this.GenerateBlockEnd(w)
}

func (this *GolangGeneratorBase) GenerateIfBegin(w io.Writer, format string, args ...interface{}) {
	this.Fprint(w, "if ")
	fmt.Fprintf(w, format, args...)
	this.GenerateBlockBegin(w, this.Config.Indent.For)
}

func (this *GolangGeneratorBase) GenerateIfEnd(w io.Writer) {
	this.GenerateBlockEnd(w)
}

func (this *GolangGeneratorBase) GenerateBlockBegin(w io.Writer, indent int) {
	fmt.Fprintln(w, " {")
	this.EnterIndent(indent)
}

func (this *GolangGeneratorBase) GenerateBlockEnd(w io.Writer) {
	this.Exit()
	this.Fprintln(w, "}")
}

func (this *GolangGeneratorBase) GetTypeNameBySize(typeSize int) string {
	switch typeSize {
	case 1:
		return "byte"
	case 2:
		return "uint16"
	case 8:
		return "uint64"
	default:
		return "uint32"
	}
}
