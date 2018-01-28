package golang

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/lioneagle/abnf/src/basic"
	"github.com/lioneagle/abnf/src/gen/charset_gen"

	"github.com/lioneagle/goutil/src/logger"
)

type CharsetTableGeneratorForGolang struct {
}

func NewCharsetTableGeneratorForGolang() *CharsetTableGeneratorForGolang {
	return &CharsetTableGeneratorForGolang{}
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

	fmt.Fprintf(file, "package %s\r\n\r\n", config.PackageName)

	if charsets != nil && len(charsets.Charsets) > 0 {
		if config.UseBit {
			fmt.Fprint(file, "/*---------------- mask definition ----------------*/\r\n")
			this.GenerateMask(config, charsets, file)
			fmt.Fprint(file, "\r\n")
		}

		fmt.Fprint(file, "/*---------------- action definition ----------------*/\r\n")
		this.GenerateAction(config, charsets, file)
		fmt.Fprint(file, "\r\n")

		fmt.Fprint(file, "/*---------------- var declaration ----------------*/\r\n")
		this.GenerateVarDeclaration(config, charsets, file)
		fmt.Fprint(file, "\r\n")
	}

	this.GenerateVarVarDefinition(config, charsets, file)
	fmt.Fprint(file, "\r\n")
}

func (this *CharsetTableGeneratorForGolang) GenerateMask(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, w io.Writer) {
	fmt.Fprintf(w, "const (\r\n")
	format := fmt.Sprintf("%s = 0x%%0%dx", getVarTypeName(config), config.VarTypeSize*2)
	for _, v := range charsets.Charsets {
		maskName := v.GetMaskName(config)

		fmt.Fprintf(w, "    %s", maskName)
		basic.PrintIndent(w, charsets.MaskNameMaxLen+4-len(maskName))
		fmt.Fprintf(w, format, v.MaskValue)
		fmt.Fprint(w, "\r\n")
	}
	fmt.Fprintf(w, ")\r\n")
}

func (this *CharsetTableGeneratorForGolang) GenerateAction(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, w io.Writer) {
	//config.ActionFirstLower = false
	for _, v := range charsets.Charsets {
		actionName := v.GetActionName(config)

		fmt.Fprintf(w, "func %s(ch byte)", actionName)
		basic.PrintIndent(w, charsets.ActionNameMaxLen+2-len(actionName))
		fmt.Fprintf(w, "{ return (%s%d[ch]", config.VarName, v.VarIndex)

		if config.UseBit {
			fmt.Fprintf(w, " & %s", v.GetMaskName(config))
		}
		fmt.Fprint(w, ") != 0 }\r\n")
	}
}

func (this *CharsetTableGeneratorForGolang) GenerateVarDeclaration(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, w io.Writer) {
}

func (this *CharsetTableGeneratorForGolang) GenerateVarVarDefinition(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, w io.Writer) {
	varTypeName := getVarTypeName(config)
	format := fmt.Sprintf("    0x%%0%dx,  /* position %%03d", config.VarTypeSize*2)

	for i := 0; i < len(charsets.Vars); i++ {
		v := &charsets.Vars[i]
		fmt.Fprintf(w, "var %s%d = [256]%s{\r\n", config.VarName, i, varTypeName)
		for j := 0; j < 256; j++ {
			ch := v.Data[j]
			fmt.Fprintf(w, format, ch, j)
			if strconv.IsPrint(rune(j)) && j <= '~' {
				fmt.Fprintf(w, "  '%c'", j)
			}
			fmt.Fprintf(w, " */\r\n")
		}
		fmt.Fprintf(w, "}\r\n\r\n")
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
