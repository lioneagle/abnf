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

func TestCGeneratorBaseGenerateVar(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_output`)

	name := "c_genertator_base_gen_var"

	standard_file := filepath.FromSlash(standard_path + "/" + name + ".c")
	output_file := filepath.FromSlash(output_path + "/" + name + ".c")

	config := gen.NewCConfigBase()

	generator := NewCGeneratorBase()
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

	generator.GenerateVarDeclare(outputFile, v1, 1, 0)

	config.VarUseSingleLineComment = false

	generator.GenerateVarDefine(outputFile, v1, 1, 0)

	generator.GenerateVarParam(outputFile, v1)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestCGeneratorBaseGenerateParamList(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_output`)

	name := "c_genertator_base_gen_param_list"

	standard_file := filepath.FromSlash(standard_path + "/" + name + ".c")
	output_file := filepath.FromSlash(output_path + "/" + name + ".c")

	config := gen.NewCConfigBase()

	generator := NewCGeneratorBase()
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

func TestCGeneratorBaseGenerateSwitch(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_output`)

	name := "c_genertator_base_gen_switch"

	standard_file := filepath.FromSlash(standard_path + "/" + name + ".c")
	output_file := filepath.FromSlash(output_path + "/" + name + ".c")

	config := gen.NewCConfigBase()

	generator := NewCGeneratorBase()
	generator.Config = config

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	generator.GenerateSwitchBegin(outputFile, "kind")
	generator.Fprintfln(outputFile, "// generate switch, left brace at next line")
	generator.GenerateCaseBegin(outputFile, "1")
	generator.Fprintfln(outputFile, "break;")
	generator.GenerateCaseEnd(outputFile)
	generator.GenerateCaseBegin(outputFile, "2")
	generator.Fprintfln(outputFile, "break;")
	generator.GenerateCaseEnd(outputFile)
	generator.GenerateSwitchEnd(outputFile)

	config.BraceAtNextLine = false

	generator.GenerateSwitchBegin(outputFile, "kind")
	generator.Fprintfln(outputFile, "// generate switch, left brace at same line")
	generator.GenerateCaseBegin(outputFile, "1")
	generator.Fprintfln(outputFile, "break;")
	generator.GenerateCaseEnd(outputFile)
	generator.GenerateCaseBegin(outputFile, "2")
	generator.Fprintfln(outputFile, "break;")
	generator.GenerateCaseEnd(outputFile)
	generator.GenerateSwitchEnd(outputFile)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestCGeneratorBaseGenerateFor(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_output`)

	name := "c_genertator_base_gen_for"

	standard_file := filepath.FromSlash(standard_path + "/" + name + ".c")
	output_file := filepath.FromSlash(output_path + "/" + name + ".c")

	config := gen.NewCConfigBase()

	generator := NewCGeneratorBase()
	generator.Config = config

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	generator.GenerateForBegin(outputFile, "i = 0; i < 100; i++")
	generator.Fprintfln(outputFile, "// generate for, left brace at next line")
	generator.GenerateForEnd(outputFile)

	config.BraceAtNextLine = false

	generator.GenerateForBegin(outputFile, "i = 0; i < 100; i++")
	generator.Fprintfln(outputFile, "// generate for, left brace at same line")
	generator.GenerateForEnd(outputFile)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestCGeneratorBaseGenerateWhile(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_output`)

	name := "c_genertator_base_gen_while"

	standard_file := filepath.FromSlash(standard_path + "/" + name + ".c")
	output_file := filepath.FromSlash(output_path + "/" + name + ".c")

	config := gen.NewCConfigBase()

	generator := NewCGeneratorBase()
	generator.Config = config

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	generator.GenerateWhileBegin(outputFile, "i < 100")
	generator.Fprintfln(outputFile, "// generate while, left brace at next line")
	generator.GenerateWhileEnd(outputFile)

	config.BraceAtNextLine = false

	generator.GenerateWhileBegin(outputFile, "i < 100")
	generator.Fprintfln(outputFile, "// generate while, left brace at same line")
	generator.GenerateWhileEnd(outputFile)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestCGeneratorBaseGenerateDoWhile(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_output`)

	name := "c_genertator_base_gen_do_while"

	standard_file := filepath.FromSlash(standard_path + "/" + name + ".c")
	output_file := filepath.FromSlash(output_path + "/" + name + ".c")

	config := gen.NewCConfigBase()

	generator := NewCGeneratorBase()
	generator.Config = config

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	generator.GenerateDoWhileBegin(outputFile)
	generator.Fprintfln(outputFile, "// generate do_while, left brace at next line")
	generator.GenerateDoWhileEnd(outputFile, "i < 100")

	config.BraceAtNextLine = false

	generator.GenerateDoWhileBegin(outputFile)
	generator.Fprintfln(outputFile, "// generate do_while, left brace at same line")
	generator.GenerateDoWhileEnd(outputFile, "i < 100")

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestCGeneratorBaseGenerateIf(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_output`)

	name := "c_genertator_base_gen_if"

	standard_file := filepath.FromSlash(standard_path + "/" + name + ".c")
	output_file := filepath.FromSlash(output_path + "/" + name + ".c")

	config := gen.NewCConfigBase()

	generator := NewCGeneratorBase()
	generator.Config = config

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	generator.GenerateIfBegin(outputFile, "i < 100")
	generator.Fprintfln(outputFile, "// generate if, left brace at next line")
	generator.GenerateIfEnd(outputFile)

	config.BraceAtNextLine = false

	generator.GenerateIfBegin(outputFile, "i < 100")
	generator.Fprintfln(outputFile, "// generate if, left brace at same line")
	generator.GenerateIfEnd(outputFile)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestCGeneratorBaseGenerateFunction(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_output`)

	name := "c_genertator_base_gen_func"

	standard_file := filepath.FromSlash(standard_path + "/" + name + ".c")
	output_file := filepath.FromSlash(output_path + "/" + name + ".c")

	config := gen.NewCConfigBase()

	generator := NewCGeneratorBase()
	generator.Config = config

	outputFile, err := os.Create(output_file)
	if err != nil {
		logger.Error("cannot open file %s", output_file)
		return
	}
	defer outputFile.Close()

	v1 := gen.NewVar()
	v1.Name = "context"
	v1.TypeName = "Context*"
	v1.InitValue = ""
	v1.Comment = "context for parsing"

	v2 := gen.NewVar()
	v2.Name = "src"
	v2.TypeName = "char const*"
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

	generator.GenerateFunctionDeclare(outputFile, func1)
	generator.PrintReturn(outputFile)

	config.ParamsInOneLine = false
	config.MultiLineCommentDecorate = true
	generator.GenerateFunctionDefine(outputFile, func1)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestCGeneratorBaseGenerateEnum(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_output`)

	name := "c_genertator_base_gen_enum"

	standard_file := filepath.FromSlash(standard_path + "/" + name + ".c")
	output_file := filepath.FromSlash(output_path + "/" + name + ".c")

	config := gen.NewCConfigBase()

	generator := NewCGeneratorBase()
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

	config.Indent.Assign = 2
	config.Indent.Comment = 2

	generator.GenerateEnumDefine(outputFile, enum)

	test.EXPECT_TRUE(t, file.FileEqual(standard_file, output_file), "file "+filepath.Base(standard_file)+" not equal")
}

func TestCGeneratorBaseGenerateStruct(t *testing.T) {
	standard_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_standard`)
	output_path := filepath.FromSlash(os.Args[len(os.Args)-1] + `/test_data/test_output`)

	name := "c_genertator_base_gen_struct"

	standard_file := filepath.FromSlash(standard_path + "/" + name + ".c")
	output_file := filepath.FromSlash(output_path + "/" + name + ".c")

	config := gen.NewCConfigBase()

	generator := NewCGeneratorBase()
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
