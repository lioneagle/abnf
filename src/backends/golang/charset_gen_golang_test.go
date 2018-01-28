package golang

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/lioneagle/abnf/src/backends"
	"github.com/lioneagle/abnf/src/gen/charset_gen"

	"github.com/lioneagle/goutil/src/file"
	"github.com/lioneagle/goutil/src/test"
)

func TestCharsetTableGeneratorForGolang_getVarTypeName(t *testing.T) {
	testdata := []struct {
		varTypeName string
		varSize     int
		wanted      string
	}{
		{"", 1, "byte"},
		{"", 2, "uint16"},
		{"", 4, "uint32"},
		{"", 8, "uint64"},

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

func TestCharsetTableGeneratorForGolang_byte_bit(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_output`)

	name := "ps_sip_charsets_byte_bit"

	standard_go := filepath.FromSlash(standard_path + "/" + name + ".go")
	output_go := filepath.FromSlash(output_path + "/" + name + ".go")

	config := backends.BuildCharsetGenConfigForTest()
	config.PackageName = "sip_charset"

	charsets := backends.BuildCharsetTableForTest(config)

	gen_c := NewCharsetTableGeneratorForGolang()
	gen_c.GenerateFile(config, charsets, name, output_path)

	test.EXPECT_TRUE(t, file.FileEqual(standard_go, output_go), "file "+filepath.Base(standard_go)+" not equal")
}

func TestCharsetTableGeneratorForGolang_byte_no_bit(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_output`)

	name := "ps_sip_charsets_byte_no_bit"

	standard_go := filepath.FromSlash(standard_path + "/" + name + ".go")
	output_go := filepath.FromSlash(output_path + "/" + name + ".go")

	config := backends.BuildCharsetGenConfigForTest()
	config.UseBit = false
	config.PackageName = "sip_charset"

	charsets := backends.BuildCharsetTableForTest(config)

	gen_c := NewCharsetTableGeneratorForGolang()
	gen_c.GenerateFile(config, charsets, name, output_path)

	test.EXPECT_TRUE(t, file.FileEqual(standard_go, output_go), "file "+filepath.Base(standard_go)+" not equal")
}

func TestCharsetTableGeneratorForGolang_dword_bit(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_output`)

	name := "ps_sip_charsets_dword_bit"

	standard_go := filepath.FromSlash(standard_path + "/" + name + ".go")
	output_go := filepath.FromSlash(output_path + "/" + name + ".go")

	config := backends.BuildCharsetGenConfigForTest()
	config.VarTypeName = "PS_DWORD"
	config.SetVarTypeSize(4)
	config.PackageName = "sip_charset"

	charsets := backends.BuildCharsetTableForTest(config)

	gen_c := NewCharsetTableGeneratorForGolang()
	gen_c.GenerateFile(config, charsets, name, output_path)

	test.EXPECT_TRUE(t, file.FileEqual(standard_go, output_go), "file "+filepath.Base(standard_go)+" not equal")
}