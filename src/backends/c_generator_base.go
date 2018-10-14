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

type TypeInfo struct {
	Name string
}

type CGeneratorBase struct {
	chars.Indent
	Config *gen.CConfigBase
}

func NewCGeneratorBase() *CGeneratorBase {
	gen := &CGeneratorBase{}
	gen.Indent.Init(0, 4)
	return gen
}

func (this *CGeneratorBase) GenerateFunctionDeclare(w io.Writer, f *gen.Function) {
	this.GenerateFunctionDefine(w, f)
	fmt.Fprint(w, ";")
}

func (this *CGeneratorBase) GenerateFunctionDefine(w io.Writer, f *gen.Function) {
	if len(f.Comment) > 0 {
		this.GenerateMultiLineComment(w, f.Comment)
	}

	this.Fprintf(w, "%s  %s(", f.ReturnType, f.Name)
	this.GenerateParamList(w, f.Params)
	fmt.Fprint(w, ")")
}

func (this *CGeneratorBase) GenerateVarDeclare(w io.Writer, v *gen.Var, typeInterval, nameInterval int) {
	this.Fprintf(w, v.TypeName)
	basic.PrintIndent(w, typeInterval)
	fmt.Fprintf(w, "%s;", v.Name)

	if len(v.Comment) > 0 {
		basic.PrintIndent(w, nameInterval+this.Config.Indent.Comment)
		this.GenerateSingleLineCommentWithoutIndent(w, v.Comment)
	} else {
		this.PrintReturn(w)
	}
}

func (this *CGeneratorBase) GenerateVarDefine(w io.Writer, v *gen.Var, typeInterval, nameInterval int) {
	this.Fprint(w, v.TypeName)
	basic.PrintIndent(w, typeInterval)
	fmt.Fprint(w, v.Name)
	if len(v.InitValue) > 0 {
		basic.PrintIndent(w, nameInterval+this.Config.Indent.Assign)
		fmt.Fprintf(w, "= %s;", v.InitValue)
	} else {
		fmt.Fprint(w, ";")

	}

	if len(v.Comment) > 0 {
		basic.PrintIndent(w, nameInterval+this.Config.Indent.Comment)
		this.GenerateSingleLineCommentWithoutIndent(w, v.Comment)
	} else {
		this.PrintReturn(w)
	}
}

func (this *CGeneratorBase) GenerateVarEnum(w io.Writer, v *gen.Var, nameInterval, valueInterval int) {
	this.Fprint(w, v.Name)
	if len(v.InitValue) > 0 {
		basic.PrintIndent(w, nameInterval+this.Config.Indent.Assign)
		fmt.Fprintf(w, "= %s,", v.InitValue)
	} else {
		fmt.Fprint(w, ",")
	}

	if len(v.Comment) > 0 {
		basic.PrintIndent(w, valueInterval+this.Config.Indent.Comment)
		this.GenerateSingleLineCommentWithoutIndent(w, v.Comment)
	}
}

func (this *CGeneratorBase) GenerateVarParam(w io.Writer, v *gen.Var) {
	fmt.Fprintf(w, "%s %s", v.TypeName, v.Name)
}

func (this *CGeneratorBase) GenerateParamList(w io.Writer, params *gen.ParamList) {
	if len(params.Params) <= 0 {
		return
	}

	this.GenerateVarParam(w, params.Params[0])

	if this.Config.ParamsInOneLine {
		for i := 1; i < len(params.Params); i++ {
			fmt.Fprintf(w, ", %s %s", params.Params[i].TypeName, params.Params[i].Name)
		}
	} else {
		this.EnterIndent(this.Config.Indent.FuncParam)
		for i := 1; i < len(params.Params); i++ {
			fmt.Fprint(w, ",")
			this.PrintReturn(w)
			this.Fprintf(w, "%s %s", params.Params[i].TypeName, params.Params[i].Name)
		}
		this.Exit()
	}
}

func (this *CGeneratorBase) GenerateMultiLineComment(w io.Writer, comment string) {
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

func (this *CGeneratorBase) GenerateSingleLineComment(w io.Writer, comment string) {
	if this.Config.VarUseSingleLineComment {
		this.Fprintfln(w, "// %s", comment)
	} else {
		this.Fprintfln(w, "/* %s */", comment)
	}
}

func (this *CGeneratorBase) GenerateSingleLineCommentWithoutIndent(w io.Writer, comment string) {
	if this.Config.VarUseSingleLineComment {
		fmt.Fprintf(w, "// %s", comment)
	} else {
		fmt.Fprintf(w, "/* %s */", comment)
	}
	this.PrintReturn(w)
}

func (this *CGeneratorBase) GenerateStructDefine(w io.Writer, s *gen.Struct) {
	this.Fprintf(w, "typedef struct tag_%s", s.Name)
	this.GenerateBlockBegin(w, this.Config.Indent.Struct)

	maxTypeNameLen := s.Fields.GetMaxTypeNameLen()
	maxNameLen := s.Fields.GetMaxNameLen()

	for _, v := range s.Fields.Params {
		this.GenerateVarDeclare(w, v, maxTypeNameLen-len(v.TypeName), maxNameLen-len(v.Name))
	}
	this.Exit()
	this.Fprintfln(w, "}%s;", s.Name)
}

func (this *CGeneratorBase) GenerateEnumDefine(w io.Writer, e *gen.Enum) {
	this.Fprintf(w, "typedef enum tag_%s", e.Name)
	this.GenerateBlockBegin(w, this.Config.Indent.Enum)

	maxNameLen := e.Enums.GetMaxNameLen()
	maxValueLen := e.Enums.GetMaxValueLen()

	for _, v := range e.Enums.Params {
		this.GenerateVarEnum(w, v, maxNameLen-len(v.Name), maxValueLen-len(v.InitValue))
	}
	this.Exit()
	this.Fprintfln(w, "}%s;", e.Name)
}

func (this *CGeneratorBase) GenerateSwitchBegin(w io.Writer, format string, args ...interface{}) {
	this.Fprint(w, "switch (")
	fmt.Fprintf(w, format, args...)
	fmt.Fprint(w, ")")

	this.GenerateBlockBegin(w, this.Config.Indent.Block)
}

func (this *CGeneratorBase) GenerateSwitchEnd(w io.Writer) {
	this.GenerateBlockEnd(w)
}

func (this *CGeneratorBase) GenerateCaseBegin(w io.Writer, format string, args ...interface{}) {
	this.Fprint(w, "case ")
	fmt.Fprintf(w, format, args...)
	fmt.Fprint(w, ":")

	this.GenerateBlockBegin(w, this.Config.Indent.Block)
}

func (this *CGeneratorBase) GenerateCaseEnd(w io.Writer) {
	this.GenerateBlockEnd(w)
}

func (this *CGeneratorBase) GenerateWhileBegin(w io.Writer, format string, args ...interface{}) {
	this.Fprint(w, "while (")
	fmt.Fprintf(w, format, args...)
	fmt.Fprint(w, ")")
	this.GenerateBlockBegin(w, this.Config.Indent.While)
}

func (this *CGeneratorBase) GenerateWhileEnd(w io.Writer) {
	this.GenerateBlockEnd(w)
}

func (this *CGeneratorBase) GenerateDoWhileBegin(w io.Writer) {
	this.Fprintf(w, "%s", "do")
	this.GenerateBlockBegin(w, this.Config.Indent.While)
}

func (this *CGeneratorBase) GenerateDoWhileEnd(w io.Writer, format string, args ...interface{}) {
	this.Exit()
	this.Fprint(w, "}while(")
	fmt.Fprintf(w, format, args...)
	fmt.Fprintln(w, ");")
}

func (this *CGeneratorBase) GenerateForBegin(w io.Writer, format string, args ...interface{}) {
	this.Fprint(w, "for (")
	fmt.Fprintf(w, format, args...)
	fmt.Fprint(w, ")")
	this.GenerateBlockBegin(w, this.Config.Indent.For)
}

func (this *CGeneratorBase) GenerateForEnd(w io.Writer) {
	this.GenerateBlockEnd(w)
}

func (this *CGeneratorBase) GenerateIfBegin(w io.Writer, format string, args ...interface{}) {
	this.Fprint(w, "if (")
	fmt.Fprintf(w, format, args...)
	fmt.Fprint(w, ")")
	this.GenerateBlockBegin(w, this.Config.Indent.For)
}

func (this *CGeneratorBase) GenerateIfEnd(w io.Writer) {
	this.GenerateBlockEnd(w)
}

func (this *CGeneratorBase) GenerateBlockBegin(w io.Writer, indent int) {
	if this.Config.BraceAtNextLine {
		fmt.Fprintln(w)
		this.Fprintln(w, "{")
	} else {
		fmt.Fprintln(w, " {")
	}
	this.EnterIndent(indent)
}

func (this *CGeneratorBase) GenerateBlockEnd(w io.Writer) {
	this.Exit()
	this.Fprintln(w, "}")
}

func (this *CGeneratorBase) GetTypeNameBySize(typeSize int) string {
	switch typeSize {
	case 1:
		return "unsigned char"
	case 2:
		return "unsigned short"
	case 8:
		return "unsigned long"
	default:
		return "unsigned int"
	}
}
