package cpp

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

func TestCharsetTableGeneratorForCpp_getVarTypeName(t *testing.T) {
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

func TestCharsetTableGeneratorForCpp_byte_bit(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/src/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/src/test_data/test_output`)

	name := "ps_sip_charsets_byte_bit"

	standard_h := filepath.FromSlash(standard_path + "/" + name + ".hpp")
	standard_c := filepath.FromSlash(standard_path + "/" + name + ".cpp")
	output_h := filepath.FromSlash(output_path + "/" + name + ".hpp")
	output_c := filepath.FromSlash(output_path + "/" + name + ".cpp")

	config := backends.BuildCharsetGenConfigForTest()
	charsets := backends.BuildCharsetTableForTest(config)

	gen_c := NewCharsetTableGeneratorForCpp()
	gen_c.GenerateFile(config, charsets, name, output_path)

	test.EXPECT_TRUE(t, file.FileEqual(standard_h, output_h), "file "+filepath.Base(standard_h)+" not equal")
	test.EXPECT_TRUE(t, file.FileEqual(standard_c, output_c), "file "+filepath.Base(standard_c)+" not equal")
}

func TestCharsetTableGeneratorForCpp_byte_no_bit(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/src/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/src/test_data/test_output`)

	name := "ps_sip_charsets_byte_no_bit"

	standard_h := filepath.FromSlash(standard_path + "/" + name + ".hpp")
	standard_c := filepath.FromSlash(standard_path + "/" + name + ".cpp")
	output_h := filepath.FromSlash(output_path + "/" + name + ".hpp")
	output_c := filepath.FromSlash(output_path + "/" + name + ".cpp")

	config := backends.BuildCharsetGenConfigForTest()
	config.UseBit = false

	charsets := backends.BuildCharsetTableForTest(config)

	gen_c := NewCharsetTableGeneratorForCpp()
	gen_c.GenerateFile(config, charsets, name, output_path)

	test.EXPECT_TRUE(t, file.FileEqual(standard_h, output_h), "file "+filepath.Base(standard_h)+" not equal")
	test.EXPECT_TRUE(t, file.FileEqual(standard_c, output_c), "file "+filepath.Base(standard_c)+" not equal")
}

func TestCharsetTableGeneratorForCpp_dword_bit(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/src/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/src/test_data/test_output`)

	name := "ps_sip_charsets_dword_bit"

	standard_h := filepath.FromSlash(standard_path + "/" + name + ".hpp")
	standard_c := filepath.FromSlash(standard_path + "/" + name + ".cpp")
	output_h := filepath.FromSlash(output_path + "/" + name + ".hpp")
	output_c := filepath.FromSlash(output_path + "/" + name + ".cpp")

	config := backends.BuildCharsetGenConfigForTest()
	config.VarTypeName = "PS_DWORD"
	config.SetVarTypeSize(4)

	charsets := backends.BuildCharsetTableForTest(config)

	gen_c := NewCharsetTableGeneratorForCpp()
	gen_c.GenerateFile(config, charsets, name, output_path)

	test.EXPECT_TRUE(t, file.FileEqual(standard_h, output_h), "file "+filepath.Base(standard_h)+" not equal")
	test.EXPECT_TRUE(t, file.FileEqual(standard_c, output_c), "file "+filepath.Base(standard_c)+" not equal")
}
