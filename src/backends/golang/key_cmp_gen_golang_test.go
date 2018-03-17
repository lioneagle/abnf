package golang

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

func TestKeyCmpGeneratorForGolang(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_output`)

	name := "ps_sip_header_key_cmp"

	standard_go := filepath.FromSlash(standard_path + "/" + name + ".go")
	output_go := filepath.FromSlash(output_path + "/" + name + ".go")

	config := backends.BuildKeyGenConfigForTest()
	config.BraceAtNextLine = false
	config.CaseSensitive = true
	config.UseTabIndent = true
	config.PackageName = "sip_header"
	config.CursorName = "pos"
	config.IndexTypeName = "SipHeaderIndexType"
	config.IndexTypeSize = 4

	keys := backends.BuildKeysForTest(config)

	gen_go := NewKeyCmpGeneratorForGolang()
	gen_go.GenerateFile(config, keys, name, output_path)

	test.EXPECT_TRUE(t, file.FileEqual(standard_go, output_go), "file "+filepath.Base(standard_go)+" not equal")
}

func TestKeyCmpGeneratorForGolang_2(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_output`)

	name := "ps_sip_header_key_cmp_2"

	standard_go := filepath.FromSlash(standard_path + "/" + name + ".go")
	output_go := filepath.FromSlash(output_path + "/" + name + ".go")

	config := backends.BuildKeyGenConfigForTest()
	config.BraceAtNextLine = false
	config.CaseSensitive = false
	config.UseTabIndent = true
	config.PackageName = "sip_header"
	config.CursorName = "pos"
	config.IndexTypeName = "SipHeaderIndexType"
	config.IndexTypeSize = 4

	keys := backends.BuildKeysForTest(config)

	gen_go := NewKeyCmpGeneratorForGolang()
	gen_go.GenerateFile(config, keys, name, output_path)

	test.EXPECT_TRUE(t, file.FileEqual(standard_go, output_go), "file "+filepath.Base(standard_go)+" not equal")
}

func TestKeyCmpGeneratorForGolang_SimpleTree_1(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_output`)

	name := "ps_sip_header_key_cmp_simple_1"

	standard_go := filepath.FromSlash(standard_path + "/" + name + ".go")
	output_go := filepath.FromSlash(output_path + "/" + name + ".go")

	config := backends.BuildKeyGenConfigForTest()
	config.BraceAtNextLine = false
	config.CaseSensitive = false
	config.UseTabIndent = true
	config.PackageName = "sip_header"
	config.CursorName = "pos"
	config.IndexTypeName = "SipHeaderIndexType"
	config.BuildSimpleTree = true

	keys := backends.BuildKeysForTest(config)

	gen_go := NewKeyCmpGeneratorForGolang()
	gen_go.GenerateFile(config, keys, name, output_path)

	test.EXPECT_TRUE(t, file.FileEqual(standard_go, output_go), "file "+filepath.Base(standard_go)+" not equal")
}
