package cpp

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/lioneagle/abnf/src/basic"
	"github.com/lioneagle/abnf/src/gen/charset_gen"

	"github.com/lioneagle/goutil/src/chars"
	"github.com/lioneagle/goutil/src/logger"
)

type CharsetTableGeneratorForCpp struct {
	chars.Indent
}

func NewCharsetTableGeneratorForCpp() *CharsetTableGeneratorForCpp {
	ret := &CharsetTableGeneratorForCpp{}
	ret.Indent.Init(0, 4)
	return ret
}

func (this *CharsetTableGeneratorForCpp) GenerateFile(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, filename, path string) {

	this.generateHFile(config, charsets, filename, path)
	this.generateCFile(config, charsets, filename, path)

}

func (this *CharsetTableGeneratorForCpp) generateHFile(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, filename, path string) {

	abs_filename := filepath.FromSlash(path + "/" + filename + ".hpp")
	file, err := os.Create(abs_filename)
	if err != nil {
		logger.Error("cannot open file %s", abs_filename)
		return
	}
	defer file.Close()

	name := strings.ToUpper(filename)

	this.Fprintfln(file, "#ifndef %s_HPP", name)
	this.Fprintfln(file, "#define %s_HPP", name)
	this.Fprintln(file)

	if charsets != nil && len(charsets.Charsets) > 0 {
		if config.UseBit {
			this.Fprintln(file, "/*---------------- mask definition ----------------*/")
			this.GenerateMask(config, charsets, file)
			this.Fprintln(file)
		}

		this.Fprintln(file, "/*---------------- action declaration ----------------*/")
		this.GenerateAction(config, charsets, file)
		this.Fprintln(file)

		this.Fprintln(file, "/*---------------- var declaration ----------------*/")
		this.GenerateVarDeclaration(config, charsets, file)
		this.Fprintln(file)
	}

	this.Fprintfln(file, "#endif /* %s_HPP */", name)
	this.Fprintln(file)
}

func (this *CharsetTableGeneratorForCpp) generateCFile(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, filename, path string) {

	abs_filename := filepath.FromSlash(path + "/" + filename + ".cpp")
	file, err := os.Create(abs_filename)
	if err != nil {
		logger.Error("cannot open file %s", abs_filename)
		return
	}
	defer file.Close()

	this.Fprintfln(file, "#include \"%s\"", filename+".h")
	this.Fprintln(file)
	this.GenerateVarDefinition(config, charsets, file)
}

func (this *CharsetTableGeneratorForCpp) GenerateMask(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, w io.Writer) {

	format := fmt.Sprintf(" = 0x%%0%dx;", config.VarTypeSize*2)
	for _, v := range charsets.Charsets {
		maskName := v.GetMaskName(config)

		this.Fprintf(w, "const %s  %s", getVarTypeName(config), maskName)
		basic.PrintIndent(w, charsets.MaskNameMaxLen+4-len(maskName))
		this.Fprintfln(w, format, v.MaskValue)
	}
}

func (this *CharsetTableGeneratorForCpp) GenerateAction(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, w io.Writer) {

	for _, v := range charsets.Charsets {
		actionName := v.GetActionName(config)

		this.Fprintf(w, "inline bool %s(unsigned char ch)", actionName)
		basic.PrintIndent(w, charsets.ActionNameMaxLen+2-len(actionName))
		fmt.Fprintf(w, "{ return %s%d[ch]", config.VarName, v.VarIndex)

		if config.UseBit {
			fmt.Fprintf(w, " & %s", v.GetMaskName(config))
		}
		this.Fprintln(w, "; }")
	}
}

func (this *CharsetTableGeneratorForCpp) GenerateVarDeclaration(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, w io.Writer) {

	varTypeName := getVarTypeName(config)
	for i := 0; i < len(charsets.Vars); i++ {
		this.Fprintfln(w, "extern %s const %s%d[256];", varTypeName, config.VarName, i)
	}
}

func (this *CharsetTableGeneratorForCpp) GenerateVarDefinition(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, w io.Writer) {

	varTypeName := getVarTypeName(config)
	format := fmt.Sprintf("0x%%0%dx,  /* position %%03d", config.VarTypeSize*2)

	for i := 0; i < len(charsets.Vars); i++ {
		v := &charsets.Vars[i]
		this.Fprintfln(w, "%s const %s%d[256] =", varTypeName, config.VarName, i)
		this.Fprintln(w, "{")
		this.Enter()
		for j := 0; j < 256; j++ {
			ch := v.Data[j]
			this.Fprintf(w, format, ch, j)
			if strconv.IsPrint(rune(j)) && j <= '~' {
				fmt.Fprintf(w, "  '%c'", j)
			}
			fmt.Fprintf(w, " */")
			this.PrintReturn(w)
		}
		this.Exit()
		this.Fprintfln(w, "};")
		this.Fprintln(w)
	}
}

func getVarTypeName(config *charset_gen.Config) string {
	if len(config.VarTypeName) > 0 {
		return config.VarTypeName
	}

	switch config.VarTypeSize {
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
