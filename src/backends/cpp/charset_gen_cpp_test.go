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
	name := "ps_sip_charsets_byte_bit"

	root := os.Args[len(os.Args)-1] + "/test_data/"

	standard_hpp, output_hpp := test.GenTestFileNames(root, "test_standard", "test_output", name+".hpp")
	standard_cpp, output_cpp := test.GenTestFileNames(root, "test_standard", "test_output", name+".cpp")

	config := backends.BuildCharsetGenConfigForTest()
	charsets := backends.BuildCharsetTableForTest(config)

	gen_cpp := NewCharsetTableGeneratorForCpp()
	gen_cpp.GenerateFile(config, charsets, name, filepath.Dir(output_cpp))

	test.EXPECT_TRUE(t, file.FileEqual(standard_hpp, output_hpp), "file "+filepath.Base(standard_hpp)+" not equal")
	test.EXPECT_TRUE(t, file.FileEqual(standard_cpp, output_cpp), "file "+filepath.Base(standard_cpp)+" not equal")
}

func TestCharsetTableGeneratorForCpp_byte_no_bit(t *testing.T) {
	name := "ps_sip_charsets_byte_no_bit"

	root := os.Args[len(os.Args)-1] + "/test_data/"

	standard_hpp, output_hpp := test.GenTestFileNames(root, "test_standard", "test_output", name+".hpp")
	standard_cpp, output_cpp := test.GenTestFileNames(root, "test_standard", "test_output", name+".cpp")

	config := backends.BuildCharsetGenConfigForTest()
	config.UseBit = false

	charsets := backends.BuildCharsetTableForTest(config)

	gen_cpp := NewCharsetTableGeneratorForCpp()
	gen_cpp.GenerateFile(config, charsets, name, filepath.Dir(output_cpp))

	test.EXPECT_TRUE(t, file.FileEqual(standard_hpp, output_hpp), "file "+filepath.Base(standard_hpp)+" not equal")
	test.EXPECT_TRUE(t, file.FileEqual(standard_cpp, output_cpp), "file "+filepath.Base(standard_cpp)+" not equal")
}

func TestCharsetTableGeneratorForCpp_dword_bit(t *testing.T) {
	name := "ps_sip_charsets_dword_bit"

	root := os.Args[len(os.Args)-1] + "/test_data/"

	standard_hpp, output_hpp := test.GenTestFileNames(root, "test_standard", "test_output", name+".hpp")
	standard_cpp, output_cpp := test.GenTestFileNames(root, "test_standard", "test_output", name+".cpp")

	config := backends.BuildCharsetGenConfigForTest()
	config.VarTypeName = "PS_DWORD"
	config.SetVarTypeSize(4)

	charsets := backends.BuildCharsetTableForTest(config)

	gen_cpp := NewCharsetTableGeneratorForCpp()
	gen_cpp.GenerateFile(config, charsets, name, filepath.Dir(output_cpp))

	test.EXPECT_TRUE(t, file.FileEqual(standard_hpp, output_hpp), "file "+filepath.Base(standard_hpp)+" not equal")
	test.EXPECT_TRUE(t, file.FileEqual(standard_cpp, output_cpp), "file "+filepath.Base(standard_cpp)+" not equal")
}
