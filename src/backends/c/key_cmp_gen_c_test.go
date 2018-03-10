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
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_output`)

	name := "ps_sip_header_key_cmp"

	standard_h := filepath.FromSlash(standard_path + "/" + name + ".h")
	standard_c := filepath.FromSlash(standard_path + "/" + name + ".c")
	output_h := filepath.FromSlash(output_path + "/" + name + ".h")
	output_c := filepath.FromSlash(output_path + "/" + name + ".c")

	config := backends.BuildKeyGenConfigForTest()
	config.CaseSensitive = true

	keys := backends.BuildKeysForTest(config)

	gen_c := NewKeyCmpGeneratorForC()
	gen_c.GenerateFile(config, keys, name, output_path)

	test.EXPECT_TRUE(t, file.FileEqual(standard_h, output_h), "file "+filepath.Base(standard_h)+" not equal")
	test.EXPECT_TRUE(t, file.FileEqual(standard_c, output_c), "file "+filepath.Base(standard_c)+" not equal")
}

func TestKeyCmpGeneratorForC_2(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_output`)

	name := "ps_sip_header_key_cmp_2"

	standard_h := filepath.FromSlash(standard_path + "/" + name + ".h")
	standard_c := filepath.FromSlash(standard_path + "/" + name + ".c")
	output_h := filepath.FromSlash(output_path + "/" + name + ".h")
	output_c := filepath.FromSlash(output_path + "/" + name + ".c")

	config := backends.BuildKeyGenConfigForTest()
	config.BraceAtNextLine = false
	config.CaseSensitive = false

	keys := backends.BuildKeysForTest(config)

	gen_c := NewKeyCmpGeneratorForC()
	gen_c.GenerateFile(config, keys, name, output_path)

	test.EXPECT_TRUE(t, file.FileEqual(standard_h, output_h), "file "+filepath.Base(standard_h)+" not equal")
	test.EXPECT_TRUE(t, file.FileEqual(standard_c, output_c), "file "+filepath.Base(standard_c)+" not equal")
}
