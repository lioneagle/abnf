package c

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/lioneagle/abnf/src/charset"
	"github.com/lioneagle/abnf/src/gen/charset_gen"

	"github.com/lioneagle/goutil/src/file"
	"github.com/lioneagle/goutil/src/test"
)

func TestCharsetTableGeneratorForC_getVarTypeName(t *testing.T) {
	testdata := []struct {
		varTypeName string
		varSize     int
		wanted      string
	}{
		{"", 1, "unsigned char"},
		{"", 2, "unsigned short"},
		{"", 4, "unsigned int"},
		{"", 8, "unsigned long"},

		{"DWORD", 1, "DWORD"},
		{"DWORD", 4, "DWORD"},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			config := charset_gen.NewConfig()
			config.VarTypeName = v.varTypeName
			config.SetVarTypeSize(v.varSize)

			test.EXPECT_EQ(t, getVarTypeName(config), v.wanted, "")
		})
	}
}

func TestCharsetTableGeneratorForC_byte_bit(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/src/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/src/test_data/test_output`)

	name := "ps_sip_charsets_byte_bit"

	standard_h := filepath.FromSlash(standard_path + "/" + name + ".h")
	standard_c := filepath.FromSlash(standard_path + "/" + name + ".c")
	output_h := filepath.FromSlash(output_path + "/" + name + ".h")
	output_c := filepath.FromSlash(output_path + "/" + name + ".c")

	config := charset_gen.NewConfig()

	config.SetMaskPrefix("PS_SIP_CHARSETS")
	config.SetActionPrefix("PS_SIP")
	config.VarTypeName = "PS_BYTE"
	config.SetVarTypeSize(1)
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

	info = charset_gen.NewCharsetInfo("lower")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{'a', 'z' + 1})
	charsets.Add(info)

	info = charset_gen.NewCharsetInfo("upper")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{'A', 'Z' + 1})
	charsets.Add(info)

	info = charset_gen.NewCharsetInfo("alphanum")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{'a', 'z' + 1})
	info.Charset.UniteRange(&charset.Range{'A', 'Z' + 1})
	info.Charset.UniteRange(&charset.Range{'0', '9' + 1})
	charsets.Add(info)

	info = charset_gen.NewCharsetInfo("hex")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{'a', 'f' + 1})
	info.Charset.UniteRange(&charset.Range{'A', 'F' + 1})
	info.Charset.UniteRange(&charset.Range{'0', '9' + 1})
	charsets.Add(info)

	info = charset_gen.NewCharsetInfo("lower-hex-alpha")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{'a', 'f' + 1})
	charsets.Add(info)

	info = charset_gen.NewCharsetInfo("upper-hex-alpha")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{'A', 'F' + 1})
	charsets.Add(info)

	info = charset_gen.NewCharsetInfo("wsp")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{' ', ' ' + 1})
	info.Charset.UniteRange(&charset.Range{'\t', '\t' + 1})
	charsets.Add(info)

	charsets.Calc(config)

	gen_c := NewCharsetTableGeneratorForC()
	gen_c.GenerateFile(config, charsets, name, output_path)

	test.EXPECT_TRUE(t, file.FileEqual(standard_h, output_h), "file "+filepath.Base(standard_h)+" not equal")
	test.EXPECT_TRUE(t, file.FileEqual(standard_c, output_c), "file "+filepath.Base(standard_c)+" not equal")
}

func TestCharsetTableGeneratorForC_byte_no_bit(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/src/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/src/test_data/test_output`)

	name := "ps_sip_charsets_byte_no_bit"

	standard_h := filepath.FromSlash(standard_path + "/" + name + ".h")
	standard_c := filepath.FromSlash(standard_path + "/" + name + ".c")
	output_h := filepath.FromSlash(output_path + "/" + name + ".h")
	output_c := filepath.FromSlash(output_path + "/" + name + ".c")

	config := charset_gen.NewConfig()

	config.SetMaskPrefix("PS_SIP_CHARSETS")
	config.SetActionPrefix("PS_SIP")
	config.VarTypeName = "PS_BYTE"
	config.SetVarTypeSize(1)
	config.SetVarName("g_sipCharsets")
	config.ActionFirstLower = true
	config.UseBit = false

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

	info = charset_gen.NewCharsetInfo("lower")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{'a', 'z' + 1})
	charsets.Add(info)

	info = charset_gen.NewCharsetInfo("upper")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{'A', 'Z' + 1})
	charsets.Add(info)

	info = charset_gen.NewCharsetInfo("alphanum")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{'a', 'z' + 1})
	info.Charset.UniteRange(&charset.Range{'A', 'Z' + 1})
	info.Charset.UniteRange(&charset.Range{'0', '9' + 1})
	charsets.Add(info)

	info = charset_gen.NewCharsetInfo("hex")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{'a', 'f' + 1})
	info.Charset.UniteRange(&charset.Range{'A', 'F' + 1})
	info.Charset.UniteRange(&charset.Range{'0', '9' + 1})
	charsets.Add(info)

	info = charset_gen.NewCharsetInfo("lower-hex-alpha")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{'a', 'f' + 1})
	charsets.Add(info)

	info = charset_gen.NewCharsetInfo("upper-hex-alpha")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{'A', 'F' + 1})
	charsets.Add(info)

	info = charset_gen.NewCharsetInfo("wsp")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{' ', ' ' + 1})
	info.Charset.UniteRange(&charset.Range{'\t', '\t' + 1})
	charsets.Add(info)

	charsets.Calc(config)

	gen_c := NewCharsetTableGeneratorForC()
	gen_c.GenerateFile(config, charsets, name, output_path)

	test.EXPECT_TRUE(t, file.FileEqual(standard_h, output_h), "file "+filepath.Base(standard_h)+" not equal")
	test.EXPECT_TRUE(t, file.FileEqual(standard_c, output_c), "file "+filepath.Base(standard_c)+" not equal")
}

func TestCharsetTableGeneratorForC_dword_bit(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/src/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/src/test_data/test_output`)

	name := "ps_sip_charsets_dword_bit"

	standard_h := filepath.FromSlash(standard_path + "/" + name + ".h")
	standard_c := filepath.FromSlash(standard_path + "/" + name + ".c")
	output_h := filepath.FromSlash(output_path + "/" + name + ".h")
	output_c := filepath.FromSlash(output_path + "/" + name + ".c")

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

	info = charset_gen.NewCharsetInfo("lower")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{'a', 'z' + 1})
	charsets.Add(info)

	info = charset_gen.NewCharsetInfo("upper")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{'A', 'Z' + 1})
	charsets.Add(info)

	info = charset_gen.NewCharsetInfo("alphanum")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{'a', 'z' + 1})
	info.Charset.UniteRange(&charset.Range{'A', 'Z' + 1})
	info.Charset.UniteRange(&charset.Range{'0', '9' + 1})
	charsets.Add(info)

	info = charset_gen.NewCharsetInfo("hex")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{'a', 'f' + 1})
	info.Charset.UniteRange(&charset.Range{'A', 'F' + 1})
	info.Charset.UniteRange(&charset.Range{'0', '9' + 1})
	charsets.Add(info)

	info = charset_gen.NewCharsetInfo("lower-hex-alpha")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{'a', 'f' + 1})
	charsets.Add(info)

	info = charset_gen.NewCharsetInfo("upper-hex-alpha")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{'A', 'F' + 1})
	charsets.Add(info)

	info = charset_gen.NewCharsetInfo("wsp")
	info.Charset = charset.NewCharset()
	info.Charset.UniteRange(&charset.Range{' ', ' ' + 1})
	info.Charset.UniteRange(&charset.Range{'\t', '\t' + 1})
	charsets.Add(info)

	charsets.Calc(config)

	gen_c := NewCharsetTableGeneratorForC()
	gen_c.GenerateFile(config, charsets, name, output_path)

	test.EXPECT_TRUE(t, file.FileEqual(standard_h, output_h), "file "+filepath.Base(standard_h)+" not equal")
	test.EXPECT_TRUE(t, file.FileEqual(standard_c, output_c), "file "+filepath.Base(standard_c)+" not equal")
}