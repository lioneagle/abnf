package golang

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/lioneagle/abnf/src/basic"
	"github.com/lioneagle/abnf/src/gen/charset_gen"

	"github.com/lioneagle/goutil/src/chars"
	"github.com/lioneagle/goutil/src/logger"
)

type CharsetTableGeneratorForGolang struct {
	chars.Indent
}

func NewCharsetTableGeneratorForGolang() *CharsetTableGeneratorForGolang {
	ret := &CharsetTableGeneratorForGolang{}
	ret.Indent.Init(0, 4)
	ret.Indent.UseTab = true
	return ret
}

func (this *CharsetTableGeneratorForGolang) GenerateFile(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, filename, path string) {

	this.generateFile(config, charsets, filename, path)
}

func (this *CharsetTableGeneratorForGolang) generateFile(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, filename, path string) {

	abs_filename := filepath.FromSlash(path + "/" + filename + ".go")
	file, err := os.Create(abs_filename)
	if err != nil {
		logger.Error("cannot open file %s", abs_filename)
		return
	}
	defer file.Close()

	this.Fprintfln(file, "package %s", config.PackageName)
	this.Fprintln(file)

	if charsets != nil && len(charsets.Charsets) > 0 {
		if config.UseBit {
			this.Fprintln(file, "/*---------------- mask definition ----------------*/")
			this.GenerateMask(config, charsets, file)
			this.Fprintln(file)
		}

		this.Fprintln(file, "/*---------------- action definition ----------------*/")
		this.GenerateAction(config, charsets, file)
		this.Fprintln(file)

		this.Fprintln(file, "/*---------------- var definition ----------------*/")
		this.GenerateVarDefinition(config, charsets, file)
		this.Fprint(file)
	}

}

func (this *CharsetTableGeneratorForGolang) GenerateMask(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, w io.Writer) {

	this.Fprintfln(w, "const (")
	this.Enter()
	format := fmt.Sprintf("%s = 0x%%0%dx", getVarTypeName(config), config.VarTypeSize*2)
	for _, v := range charsets.Charsets {
		maskName := v.GetMaskName(config)

		this.Fprintf(w, "%s", maskName)
		basic.PrintIndent(w, charsets.MaskNameMaxLen+4-len(maskName))
		fmt.Fprintf(w, format, v.MaskValue)
		this.PrintReturn(w)
	}
	this.Exit()
	this.Fprintln(w, ")")
}

func (this *CharsetTableGeneratorForGolang) GenerateAction(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, w io.Writer) {

	//config.ActionFirstLower = false
	for _, v := range charsets.Charsets {
		actionName := v.GetActionName(config)

		this.Fprintf(w, "func %s(ch byte)", actionName)
		basic.PrintIndent(w, charsets.ActionNameMaxLen+2-len(actionName))
		this.Fprintf(w, "{ return (%s%d[ch]", config.VarName, v.VarIndex)

		if config.UseBit {
			this.Fprintf(w, " & %s", v.GetMaskName(config))
		}
		this.Fprintln(w, ") != 0 }")
	}
}

func (this *CharsetTableGeneratorForGolang) GenerateVarDeclaration(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, w io.Writer) {
}

func (this *CharsetTableGeneratorForGolang) GenerateVarDefinition(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, w io.Writer) {

	varTypeName := getVarTypeName(config)
	format := fmt.Sprintf("0x%%0%dx,  /* position %%03d", config.VarTypeSize*2)

	for i := 0; i < len(charsets.Vars); i++ {
		v := &charsets.Vars[i]
		this.Fprintfln(w, "var %s%d = [256]%s{", config.VarName, i, varTypeName)
		this.Enter()
		for j := 0; j < 256; j++ {
			ch := v.Data[j]
			this.Fprintf(w, format, ch, j)
			if strconv.IsPrint(rune(j)) && j <= '~' {
				fmt.Fprintf(w, "  '%c'", j)
			}
			fmt.Fprint(w, " */")
			this.PrintReturn(w)
		}
		this.Exit()
		this.Fprintln(w, "}")
		this.Fprintln(w)
	}
}

func getVarTypeName(config *charset_gen.Config) string {
	if len(config.VarTypeName) > 0 {
		return config.VarTypeName
	}

	switch config.VarTypeSize {
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
