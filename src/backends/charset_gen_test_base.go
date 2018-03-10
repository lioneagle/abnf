package backends

import (
	"github.com/lioneagle/abnf/src/charset"
	"github.com/lioneagle/abnf/src/gen/charset_gen"
)

func BuildCharsetTableForTest(config *charset_gen.Config) *charset_gen.CharsetTable {
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
	return charsets
}

func BuildCharsetGenConfigForTest() *charset_gen.Config {
	config := charset_gen.NewConfig()

	config.SetMaskPrefix("PS_SIP_CHARSETS")
	config.SetActionPrefix("PS_SIP")
	config.VarTypeName = "PS_BYTE"
	config.SetVarTypeSize(1)
	config.SetVarName("g_sipCharsets")
	config.ActionFirstLower = true
	config.UseBit = true

	return config
}
