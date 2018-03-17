package cpp

import (
	//"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/lioneagle/abnf/src/backends"
	//"github.com/lioneagle/abnf/src/gen/key_gen"

	"github.com/lioneagle/goutil/src/file"
	"github.com/lioneagle/goutil/src/test"
)

func TestKeyCmpGeneratorForCpp(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_output`)

	name := "ps_sip_header_key_cmp"

	standard_hpp := filepath.FromSlash(standard_path + "/" + name + ".hpp")
	standard_cpp := filepath.FromSlash(standard_path + "/" + name + ".cpp")
	output_hpp := filepath.FromSlash(output_path + "/" + name + ".hpp")
	output_cpp := filepath.FromSlash(output_path + "/" + name + ".cpp")

	config := backends.BuildKeyGenConfigForTest()
	config.CaseSensitive = true

	keys := backends.BuildKeysForTest(config)

	gen_cpp := NewKeyCmpGeneratorForCpp()
	gen_cpp.GenerateFile(config, keys, name, output_path)

	test.EXPECT_TRUE(t, file.FileEqual(standard_hpp, output_hpp), "file "+filepath.Base(standard_hpp)+" not equal")
	test.EXPECT_TRUE(t, file.FileEqual(standard_cpp, output_cpp), "file "+filepath.Base(standard_cpp)+" not equal")
}

func TestKeyCmpGeneratorForCpp_2(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_output`)

	name := "ps_sip_header_key_cmp_2"

	standard_hpp := filepath.FromSlash(standard_path + "/" + name + ".hpp")
	standard_cpp := filepath.FromSlash(standard_path + "/" + name + ".cpp")
	output_hpp := filepath.FromSlash(output_path + "/" + name + ".hpp")
	output_cpp := filepath.FromSlash(output_path + "/" + name + ".cpp")

	config := backends.BuildKeyGenConfigForTest()
	config.BraceAtNextLine = false
	config.CaseSensitive = false

	keys := backends.BuildKeysForTest(config)

	gen_cpp := NewKeyCmpGeneratorForCpp()
	gen_cpp.GenerateFile(config, keys, name, output_path)

	test.EXPECT_TRUE(t, file.FileEqual(standard_hpp, output_hpp), "file "+filepath.Base(standard_hpp)+" not equal")
	test.EXPECT_TRUE(t, file.FileEqual(standard_cpp, output_cpp), "file "+filepath.Base(standard_cpp)+" not equal")
}

func TestKeyCmpGeneratorForCpp_SimpleTree_1(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_output`)

	name := "ps_sip_header_key_cmp_simple_1"

	standard_hpp := filepath.FromSlash(standard_path + "/" + name + ".hpp")
	standard_cpp := filepath.FromSlash(standard_path + "/" + name + ".cpp")
	output_hpp := filepath.FromSlash(output_path + "/" + name + ".hpp")
	output_cpp := filepath.FromSlash(output_path + "/" + name + ".cpp")

	config := backends.BuildKeyGenConfigForTest()
	config.BraceAtNextLine = false
	config.CaseSensitive = false
	config.BuildSimpleTree = true
	config.UseTabIndent = true

	keys := backends.BuildKeysForTest(config)

	gen_cpp := NewKeyCmpGeneratorForCpp()
	gen_cpp.GenerateFile(config, keys, name, output_path)

	test.EXPECT_TRUE(t, file.FileEqual(standard_hpp, output_hpp), "file "+filepath.Base(standard_hpp)+" not equal")
	test.EXPECT_TRUE(t, file.FileEqual(standard_cpp, output_cpp), "file "+filepath.Base(standard_cpp)+" not equal")
}
