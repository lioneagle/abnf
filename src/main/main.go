package main

import (
	//"charset"
	//"fmt"
	//"io"
	//"os"
	//"reflect"

	"github.com/lioneagle/abnf/src/backends/c"
	"github.com/lioneagle/abnf/src/charset"
	"github.com/lioneagle/abnf/src/gen/charset_gen"
)

type A struct {
	x int
}

func (a *A) f(int) bool {
	return false
}

func main() {
	config := charset_gen.NewConfig()

	config.SetMaskPrefix("PS_SIP_CHARSETS")
	config.SetActionPrefix("PS_SIP")
	config.VarTypeName = "PS_DWORD"
	config.SetVarTypeSize(4)
	config.SetVarName("g_sipCharsets")
	config.ActionFirstLower = true
	config.UseBit = true

	charsets := charset_gen.NewCharsetTable()

	info := charset_gen.NewCharsetInfo("digit")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{'0', '9' + 1})
	charsets.Add(info)

	info = charset_gen.NewCharsetInfo("alpha")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{'a', 'z' + 1})
	info.Charset.UniteRange(&charset.Range{'A', 'Z' + 1})
	charsets.Add(info)

	charsets.Calc(config)

	gen_c := c.NewCharsetTableGeneratorForC()
	gen_c.GenerateFile(config, charsets, "ps_sip_charsets_1", ".")

	/*
		var gen charset.CharsetGenForCpp

		gen.GenerateMask(os.Stdout, nil)

		var r1 charset.Range

		//fmt.Printf("0x%x\n", uint32(-1))

		r1 = charset.Range{1, 2}

		r1.PrintAsChar(os.Stdout).WriteString("\n")

		r1 = charset.Range{1, 6}
		r1.PrintAsChar(os.Stdout).WriteString("\n")

		r1 = charset.Range{0, 257}
		r1.Print(os.Stdout).WriteString("\n")
		r1.PrintEachChar(os.Stdout).WriteString("\n")

		fmt.Println("r1 = ", r1)
		fmt.Printf("%c\n", '\\')

		var a A

		p := a.f

		fmt.Println("type =", reflect.TypeOf(p))
	*/
}
