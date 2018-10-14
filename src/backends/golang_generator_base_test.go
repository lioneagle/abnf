package backends

import (
	//"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/lioneagle/abnf/src/gen"
	//"github.com/lioneagle/abnf/src/gen/charset_gen"

	"github.com/lioneagle/goutil/src/file"
	"github.com/lioneagle/goutil/src/logger"
	"github.com/lioneagle/goutil/src/test"
)

func TestGolangGeneratorBaseGenerateVar(t *testing.T) {
	standard_file, output_file := test.GenTestFileNames(os.Args[len(os.Args)-1]+"/test_data/", "test_standard", "test_output", "golang_genertator_base_gen_var.go")

	config := gen.NewGolangConfigBase()

	generator := NewGolangGeneratorBase()
	generator.Config = config

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	v1 := gen.NewVar()
	v1.Name = "name"
	v1.TypeName = "int"
	v1.InitValue = "100"
	v1.Comment = "name of book"

	generator.GenerateVarDeclare(outputFile, v1, 1, 1)

	config.VarUseSingleLineComment = false

	generator.GenerateVarDefine(outputFile, v1, 1, 1)

	generator.GenerateVarParam(outputFile, v1)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestGolangGeneratorBaseGenerateParamList(t *testing.T) {
	standard_file, output_file := test.GenTestFileNames(os.Args[len(os.Args)-1]+"/test_data/", "test_standard", "test_output", "golang_genertator_base_gen_param_list.go")

	config := gen.NewGolangConfigBase()

	generator := NewGolangGeneratorBase()
	generator.Config = config

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	v1 := gen.NewVar()
	v1.Name = "name"
	v1.TypeName = "int"
	v1.InitValue = "100"
	v1.Comment = "name of book"

	v2 := gen.NewVar()
	v2.Name = "value"
	v2.TypeName = "int"
	v2.InitValue = "100"
	v2.Comment = "value of book"

	v3 := gen.NewVar()
	v3.Name = "note"
	v3.TypeName = "int"
	v3.InitValue = "110"
	v3.Comment = "note of book"

	params := gen.NewParamList()
	params.Append(v1)
	params.Append(v2)
	params.Append(v3)

	generator.GenerateParamList(outputFile, params)

	config.ParamsInOneLine = false

	generator.PrintReturn(outputFile)

	generator.GenerateParamList(outputFile, params)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestGolangGeneratorBaseGenerateSwitch(t *testing.T) {
	standard_file, output_file := test.GenTestFileNames(os.Args[len(os.Args)-1]+"/test_data/", "test_standard", "test_output", "golang_genertator_base_gen_switch.go")

	config := gen.NewGolangConfigBase()

	generator := NewGolangGeneratorBase()
	generator.Config = config

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	generator.GenerateSwitchBegin(outputFile, "kind")
	generator.GenerateCaseBegin(outputFile, "1")
	generator.Fprintfln(outputFile, "break")
	generator.GenerateCaseEnd(outputFile)
	generator.GenerateCaseBegin(outputFile, "2")
	generator.Fprintfln(outputFile, "break")
	generator.GenerateCaseEnd(outputFile)
	generator.GenerateSwitchEnd(outputFile)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestGolangGeneratorBaseGenerateFor(t *testing.T) {
	standard_file, output_file := test.GenTestFileNames(os.Args[len(os.Args)-1]+"/test_data/", "test_standard", "test_output", "golang_genertator_base_gen_for.go")

	config := gen.NewGolangConfigBase()

	generator := NewGolangGeneratorBase()
	generator.Config = config

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	generator.GenerateForBegin(outputFile, "i = 0; i < 100; i++")
	generator.GenerateForEnd(outputFile)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestGolangGeneratorBaseGenerateIf(t *testing.T) {
	standard_file, output_file := test.GenTestFileNames(os.Args[len(os.Args)-1]+"/test_data/", "test_standard", "test_output", "golang_genertator_base_gen_if.go")

	config := gen.NewGolangConfigBase()

	generator := NewGolangGeneratorBase()
	generator.Config = config

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	generator.GenerateIfBegin(outputFile, "i < 100")
	generator.GenerateIfEnd(outputFile)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestGolangGeneratorBaseGenerateFunction(t *testing.T) {
	standard_file, output_file := test.GenTestFileNames(os.Args[len(os.Args)-1]+"/test_data/", "test_standard", "test_output", "golang_genertator_base_gen_func.go")
	config := gen.NewGolangConfigBase()

	generator := NewGolangGeneratorBase()
	generator.Config = config

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	v1 := gen.NewVar()
	v1.Name = "context"
	v1.TypeName = "*Context"
	v1.InitValue = ""
	v1.Comment = "context for parsing"

	v2 := gen.NewVar()
	v2.Name = "src"
	v2.TypeName = "string"
	v2.InitValue = ""
	v2.Comment = "source to parse"

	v3 := gen.NewVar()
	v3.Name = "len"
	v3.TypeName = "int"
	v3.InitValue = ""
	v3.Comment = "length of source"

	params := gen.NewParamList()
	params.Append(v1)
	params.Append(v2)
	params.Append(v3)

	func1 := gen.NewFunction()
	func1.Name = "parse"
	func1.ReturnType = "int"
	func1.Comment = `
    NAME: parse
    PARAMS: 
    context -- context for parsing
    src     -- source to parse
    len     -- length of source
    RETURN:
    length parsed
    NOTE: create for test
    `
	func1.Params = params

	generator.GenerateFunctionDefine(outputFile, func1)
	generator.GenerateBlockBegin(outputFile, config.Indent.Block)
	generator.GenerateBlockEnd(outputFile)
	generator.PrintReturn(outputFile)

	config.ParamsInOneLine = false
	config.MultiLineCommentDecorate = true
	generator.GenerateFunctionDefine(outputFile, func1)
	generator.GenerateBlockBegin(outputFile, config.Indent.Block)
	generator.GenerateBlockEnd(outputFile)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestGolangGeneratorBaseGenerateEnum(t *testing.T) {
	standard_file, output_file := test.GenTestFileNames(os.Args[len(os.Args)-1]+"/test_data/", "test_standard", "test_output", "golang_genertator_base_gen_enum.go")
	config := gen.NewGolangConfigBase()

	generator := NewGolangGeneratorBase()
	generator.Config = config

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	v1 := gen.NewVar()
	v1.Name = "name"
	v1.TypeName = "int"
	v1.InitValue = "100"
	v1.Comment = "name of book"

	v2 := gen.NewVar()
	v2.Name = "value"
	v2.TypeName = "int"
	v2.InitValue = "100"
	v2.Comment = "value of book"

	v3 := gen.NewVar()
	v3.Name = "note"
	v3.TypeName = "int"
	v3.InitValue = "10"
	v3.Comment = "note of book"

	enum := gen.NewEnum()
	enum.Name = "ATTR_TYPE"
	enum.AppendEnum(v1)
	enum.AppendEnum(v2)
	enum.AppendEnum(v3)

	generator.GenerateEnumDefine(outputFile, enum)
	generator.PrintReturn(outputFile)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestGolangGeneratorBaseGenerateStruct(t *testing.T) {
	standard_file, output_file := test.GenTestFileNames(os.Args[len(os.Args)-1]+"/test_data/", "test_standard", "test_output", "golang_genertator_base_gen_struct.go")
	config := gen.NewGolangConfigBase()

	generator := NewGolangGeneratorBase()
	generator.Config = config

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	v1 := gen.NewVar()
	v1.Name = "name"
	v1.TypeName = "int"
	v1.InitValue = "100"
	v1.Comment = "name of book"

	v2 := gen.NewVar()
	v2.Name = "value"
	v2.TypeName = "int"
	v2.InitValue = "100"
	v2.Comment = "value of book"

	v3 := gen.NewVar()
	v3.Name = "note"
	v3.TypeName = "int"
	v3.InitValue = "10"
	v3.Comment = "note of book"

	s := gen.NewStruct()
	s.Name = "ATTR_TYPE"
	s.AppendField(v1)
	s.AppendField(v2)
	s.AppendField(v3)

	generator.GenerateStructDefine(outputFile, s)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}
