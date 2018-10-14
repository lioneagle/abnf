package c

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
	name := "ps_sip_charsets_byte_bit"
	root := os.Args[len(os.Args)-1] + "/test_data/"

	standard_h, output_h := test.GenTestFileNames(root, "test_standard", "test_output", name+".h")
	standard_c, output_c := test.GenTestFileNames(root, "test_standard", "test_output", name+".c")

	config := backends.BuildCharsetGenConfigForTest()
	charsets := backends.BuildCharsetTableForTest(config)

	generator := NewCharsetTableGeneratorForC()
	generator.GenerateFile(config, charsets, name, filepath.Dir(output_c))

	test.EXPECT_TRUE(t, file.FileEqual(standard_h, output_h), "file "+filepath.Base(standard_h)+" not equal")
	test.EXPECT_TRUE(t, file.FileEqual(standard_c, output_c), "file "+filepath.Base(standard_c)+" not equal")
}

func TestCharsetTableGeneratorForC_byte_no_bit(t *testing.T) {
	name := "ps_sip_charsets_byte_no_bit"
	root := os.Args[len(os.Args)-1] + "/test_data/"

	standard_h, output_h := test.GenTestFileNames(root, "test_standard", "test_output", name+".h")
	standard_c, output_c := test.GenTestFileNames(root, "test_standard", "test_output", name+".c")

	config := backends.BuildCharsetGenConfigForTest()
	config.UseBit = false

	charsets := backends.BuildCharsetTableForTest(config)

	generator := NewCharsetTableGeneratorForC()
	generator.GenerateFile(config, charsets, name, filepath.Dir(output_c))

	test.EXPECT_TRUE(t, file.FileEqual(standard_h, output_h), "file "+filepath.Base(standard_h)+" not equal")
	test.EXPECT_TRUE(t, file.FileEqual(standard_c, output_c), "file "+filepath.Base(standard_c)+" not equal")
}

func TestCharsetTableGeneratorForC_dword_bit(t *testing.T) {
	name := "ps_sip_charsets_dword_bit"
	root := os.Args[len(os.Args)-1] + "/test_data/"

	standard_h, output_h := test.GenTestFileNames(root, "test_standard", "test_output", name+".h")
	standard_c, output_c := test.GenTestFileNames(root, "test_standard", "test_output", name+".c")

	config := backends.BuildCharsetGenConfigForTest()
	config.VarTypeName = "PS_DWORD"
	config.SetVarTypeSize(4)

	charsets := backends.BuildCharsetTableForTest(config)

	gen_c := NewCharsetTableGeneratorForC()
	gen_c.GenerateFile(config, charsets, name, filepath.Dir(output_c))

	test.EXPECT_TRUE(t, file.FileEqual(standard_h, output_h), "file "+filepath.Base(standard_h)+" not equal")
	test.EXPECT_TRUE(t, file.FileEqual(standard_c, output_c), "file "+filepath.Base(standard_c)+" not equal")
}
