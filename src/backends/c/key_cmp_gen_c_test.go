package c

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

func TestKeyCmpGeneratorForC(t *testing.T) {
	name := "ps_sip_header_key_cmp"
	root := os.Args[len(os.Args)-1] + "/test_data/"

	standard_h, output_h := test.GenTestFileNames(root, "test_standard", "test_output", name+".h")
	standard_c, output_c := test.GenTestFileNames(root, "test_standard", "test_output", name+".c")

	config := backends.BuildKeyGenConfigForTest()
	config.CaseSensitive = true

	keys := backends.BuildKeysForTest(config)

	gen_c := NewKeyCmpGeneratorForC()
	gen_c.GenerateFile(config, keys, name, filepath.Dir(output_c))

	test.EXPECT_TRUE(t, file.FileEqual(standard_h, output_h), "file "+filepath.Base(standard_h)+" not equal")
	test.EXPECT_TRUE(t, file.FileEqual(standard_c, output_c), "file "+filepath.Base(standard_c)+" not equal")
}

func TestKeyCmpGeneratorForC_2(t *testing.T) {
	name := "ps_sip_header_key_cmp_2"
	root := os.Args[len(os.Args)-1] + "/test_data/"

	standard_h, output_h := test.GenTestFileNames(root, "test_standard", "test_output", name+".h")
	standard_c, output_c := test.GenTestFileNames(root, "test_standard", "test_output", name+".c")

	config := backends.BuildKeyGenConfigForTest()
	config.BraceAtNextLine = false
	config.CaseSensitive = false

	keys := backends.BuildKeysForTest(config)

	gen_c := NewKeyCmpGeneratorForC()
	gen_c.GenerateFile(config, keys, name, filepath.Dir(output_c))

	test.EXPECT_TRUE(t, file.FileEqual(standard_h, output_h), "file "+filepath.Base(standard_h)+" not equal")
	test.EXPECT_TRUE(t, file.FileEqual(standard_c, output_c), "file "+filepath.Base(standard_c)+" not equal")
}

func TestKeyCmpGeneratorForC_SimpleTree_1(t *testing.T) {
	name := "ps_sip_header_key_cmp_simple_1"

	root := os.Args[len(os.Args)-1] + "/test_data/"

	standard_h, output_h := test.GenTestFileNames(root, "test_standard", "test_output", name+".h")
	standard_c, output_c := test.GenTestFileNames(root, "test_standard", "test_output", name+".c")

	config := backends.BuildKeyGenConfigForTest()
	config.BraceAtNextLine = false
	config.CaseSensitive = false
	config.BuildSimpleTree = true
	config.UseTabIndent = true

	keys := backends.BuildKeysForTest(config)

	gen_c := NewKeyCmpGeneratorForC()
	gen_c.GenerateFile(config, keys, name, filepath.Dir(output_c))

	test.EXPECT_TRUE(t, file.FileEqual(standard_h, output_h), "file "+filepath.Base(standard_h)+" not equal")
	test.EXPECT_TRUE(t, file.FileEqual(standard_c, output_c), "file "+filepath.Base(standard_c)+" not equal")
}
