package c

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/lioneagle/abnf/src/basic"
	"github.com/lioneagle/abnf/src/gen/charset_gen"

	"github.com/lioneagle/goutil/src/logger"
)

type CharsetTableGeneratorForC struct {
}

func NewCharsetTableGeneratorForC() *CharsetTableGeneratorForC {
	return &CharsetTableGeneratorForC{}
}

func (this *CharsetTableGeneratorForC) GenerateFile(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, filename, path string) {
	this.generateHFile(config, charsets, filename, path)
	this.generateCFile(config, charsets, filename, path)

}

func (this *CharsetTableGeneratorForC) generateHFile(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, filename, path string) {
	abs_filename := filepath.FromSlash(path + "/" + filename + ".h")
	file, err := os.Create(abs_filename)
	if err != nil {
		logger.Error("cannot open file %s", abs_filename)
		return
	}
	defer file.Close()

	name := strings.ToUpper(filename)

	fmt.Fprintf(file, "#ifndef %s_H\r\n", name)
	fmt.Fprintf(file, "#define %s_H\r\n", name)
	fmt.Fprint(file, "\r\n")

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

	fmt.Fprintf(file, "#endif /* %s_H */\r\n", name)
}

func (this *CharsetTableGeneratorForC) generateCFile(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, filename, path string) {
	abs_filename := filepath.FromSlash(path + "/" + filename + ".c")
	file, err := os.Create(abs_filename)
	if err != nil {
		logger.Error("cannot open file %s", abs_filename)
		return
	}
	defer file.Close()

	fmt.Fprintf(file, "#include \"%s\"\r\n\r\n", filename+".h")
	this.GenerateVarVarDefinition(config, charsets, file)
}

func (this *CharsetTableGeneratorForC) GenerateMask(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, w io.Writer) {
	format := fmt.Sprintf("((%s)(0x%%0%dx))", getVarTypeName(config), config.VarTypeSize*2)
	for _, v := range charsets.Charsets {
		maskName := v.GetMaskName(config)

		fmt.Fprintf(w, "#define %s", maskName)
		basic.PrintIndent(w, charsets.MaskNameMaxLen+4-len(maskName))
		fmt.Fprintf(w, format, v.MaskValue)
		fmt.Fprint(w, "\r\n")
	}
}

func (this *CharsetTableGeneratorForC) GenerateAction(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, w io.Writer) {
	for _, v := range charsets.Charsets {
		actionName := v.GetActionName(config)

		fmt.Fprintf(w, "#define %s(ch)", actionName)
		basic.PrintIndent(w, charsets.ActionNameMaxLen+4-len(actionName))
		fmt.Fprintf(w, "(%s%d[(unsigned char)(ch)]", config.VarName, v.VarIndex)

		if config.UseBit {
			fmt.Fprintf(w, " & %s", v.GetMaskName(config))
		}
		fmt.Fprint(w, ")\r\n")
	}
}

func (this *CharsetTableGeneratorForC) GenerateVarDeclaration(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, w io.Writer) {
	varTypeName := getVarTypeName(config)
	for i := 0; i < len(charsets.Vars); i++ {
		fmt.Fprintf(w, "extern %s const %s%d[256];\r\n", varTypeName, config.VarName, i)
	}
}

func (this *CharsetTableGeneratorForC) GenerateVarVarDefinition(config *charset_gen.Config,
	charsets *charset_gen.CharsetTable, w io.Writer) {
	varTypeName := getVarTypeName(config)
	format := fmt.Sprintf("    0x%%0%dx,  /* position %%03d", config.VarTypeSize*2)

	for i := 0; i < len(charsets.Vars); i++ {
		v := &charsets.Vars[i]
		fmt.Fprintf(w, "%s const %s%d[256] =\r\n{\r\n", varTypeName, config.VarName, i)
		for j := 0; j < 256; j++ {
			ch := v.Data[j]
			fmt.Fprintf(w, format, ch, j)
			if strconv.IsPrint(rune(j)) && j <= '~' {
				fmt.Fprintf(w, "  '%c'", j)
			}
			fmt.Fprintf(w, " */\r\n")
		}
		fmt.Fprintf(w, "};\r\n\r\n")
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