package cpp

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/lioneagle/abnf/src/basic"
	"github.com/lioneagle/abnf/src/gen/key_gen"
	"github.com/lioneagle/abnf/src/gen/key_gen/key_cmp_gen"
	"github.com/lioneagle/abnf/src/keys"

	"github.com/lioneagle/goutil/src/chars"
	"github.com/lioneagle/goutil/src/logger"
	"github.com/lioneagle/goutil/src/times"
)

type KeyCmpGeneratorForCpp struct {
	chars.Indent
}

func NewKeyCmpGeneratorForCpp() *KeyCmpGeneratorForCpp {
	ret := &KeyCmpGeneratorForCpp{}
	ret.Indent.Init(0, 4)
	return ret
}

func (this *KeyCmpGeneratorForCpp) GenerateFile(config *key_gen.Config,
	keys *keys.Keys, filename, path string) {

	this.Indent.UseTab = config.UseTabIndent

	timeStat := times.NewTimeStat()
	defer func() {
		timeStat.Stop()
		if config.PrintTimeUsed {
			timeStat.Fprint(config.OutputFile, "tokencmp generate files")
		}
	}()

	this.generateHFile(config, keys, filename, path)
	tree := this.buildTrieTree(config, keys)
	this.generateCFile(config, tree, filename, path)
}

func (this *KeyCmpGeneratorForCpp) buildTrieTree(config *key_gen.Config, keys *keys.Keys) *key_cmp_gen.TrieTree {
	timeStat := times.NewTimeStat()
	defer func() {
		timeStat.Stop()
		if config.PrintTimeUsed {
			timeStat.Fprint(config.OutputFile, "tokencmp build trie tree")
		}
	}()

	return key_cmp_gen.BuildTrieTreeFromKeys(config, keys)
}

func (this *KeyCmpGeneratorForCpp) generateHFile(config *key_gen.Config,
	keys *keys.Keys, filename, path string) {

	timeStat := times.NewTimeStat()
	defer func() {
		timeStat.Stop()
		if config.PrintTimeUsed {
			timeStat.Fprint(config.OutputFile, "tokencmp generate hpp file")
		}
	}()

	abs_filename := filepath.FromSlash(path + "/" + filename + ".hpp")
	file, err := os.Create(abs_filename)
	if err != nil {
		logger.Error("cannot open file %s", abs_filename)
		return
	}
	defer file.Close()

	name := strings.ToUpper(filename)

	this.Fprintfln(file, "#ifndef %s_HPP", name)
	this.Fprintfln(file, "#define %s_HPP", name)
	this.Fprintln(file)

	if keys != nil && len(keys.Data) > 0 {

		this.Fprintln(file, "/*---------------- index definition ----------------*/")
		this.GenerateIndex(config, keys, file)
		this.Fprintln(file)

		this.Fprintln(file, "/*---------------- action declaration ----------------*/")
		this.GenerateActionDeclaration(config, file)
		this.Fprintln(file, ";")
		this.Fprintln(file)
	}

	this.Fprintfln(file, "#endif /* %s_HPP */", name)
}

func (this *KeyCmpGeneratorForCpp) generateCFile(config *key_gen.Config,
	tree *key_cmp_gen.TrieTree, filename, path string) {

	timeStat := times.NewTimeStat()
	defer func() {
		timeStat.Stop()
		if config.PrintTimeUsed {
			timeStat.Fprint(config.OutputFile, "tokencmp generate cpp file")
		}
	}()

	abs_filename := filepath.FromSlash(path + "/" + filename + ".cpp")
	file, err := os.Create(abs_filename)
	if err != nil {
		logger.Error("cannot open file %s", abs_filename)
		return
	}
	defer file.Close()

	this.Fprintfln(file, "#include \"%s\"", filename+".hpp")
	this.Fprintln(file)
	this.GenerateActionDefinition(config, tree, file)
}

func (this *KeyCmpGeneratorForCpp) GenerateIndex(config *key_gen.Config,
	keys *keys.Keys, w io.Writer) {

	MaxIndexName := keys.GetMaxIndexNameLen()

	for _, v := range keys.Indices {
		this.Fprintf(w, "const %s  %s", getIndexTypeName(config), v.Name)
		basic.PrintIndent(w, MaxIndexName+4-len(v.Name))
		this.Fprintfln(w, "= %d;", v.Value)
	}
}

func (this *KeyCmpGeneratorForCpp) GenerateActionDeclaration(config *key_gen.Config, w io.Writer) {
	indexType := getIndexTypeName(config)
	srcType := getSrcTypeName(config)

	this.Fprintf(w, "%s %s(%s* %s, %s end)", indexType, config.ActionName, srcType, config.SrcName, srcType)
}

func (this *KeyCmpGeneratorForCpp) GenerateActionDefinition(config *key_gen.Config,
	tree *key_cmp_gen.TrieTree, w io.Writer) {

	srcName := config.SrcName
	srcType := getSrcTypeName(config)

	this.GenerateActionDeclaration(config, w)

	this.GenerateBlockBegin(config, w, "")

	this.Fprintfln(w, "%s  %s = *%s;", srcType, config.CursorName, srcName)
	this.Fprintfln(w, "")

	this.Fprintf(w, "if (%s == NULL || %s >= end)", config.CursorName, config.CursorName)
	this.generateLeftBrace(config, w, config.IndentOfBlock)
	this.Fprintfln(w, "return %s;", config.UnknownIndexName)
	this.generateRightBrace(config, w)
	this.Fprintln(w)

	this.GenerateActionCode(config, tree, w, 0)

	this.Fprintln(w)
	this.Fprintfln(w, "return %s;", config.UnknownIndexName)

	this.GenerateBlockEnd(config, w)
}

func (this *KeyCmpGeneratorForCpp) GenerateActionCode(config *key_gen.Config,
	tree *key_cmp_gen.TrieTree, w io.Writer, depth int) {

	srcName := config.SrcName
	seperatorName := config.SeperatorCharsetName
	cursorName := config.CursorName

	branch := tree.FindFinal(config)

	if branch != nil {
		if config.SeperatorEnabled {
			this.Fprintf(w, "if ((%s < end) && %s(*%s))", cursorName, seperatorName, cursorName)
		} else {
			this.Fprintf(w, "if (%s >= end)", cursorName)
		}
		this.generateLeftBrace(config, w, config.IndentOfBlock)

		this.Fprintfln(w, "*%s = %s;", srcName, cursorName)
		this.Fprintfln(w, "return %s;", branch.Next.Key.Index.Name)
		this.generateRightBrace(config, w)

		if depth == 0 && len(tree.Branches) == 1 {
			this.GenerateBlockBegin(config, w, "")
			this.Fprintfln(w, "*%s = %s;", srcName, cursorName)
			this.Fprintfln(w, "return %s;", config.UnknownIndexName)
			this.GenerateBlockEnd(config, w)
			return
		}
	}

	if tree.NonFinalBranchNum() == 1 {
		this.GenerateActionCodeForSingleBranch(config, tree, w, depth)
	} else if tree.NonFinalBranchNum() > 1 {
		this.GenerateActionCodeForMultiBranch(config, tree, w, depth)
	}
}

func (this *KeyCmpGeneratorForCpp) GenerateActionCodeForMultiBranch(config *key_gen.Config,
	tree *key_cmp_gen.TrieTree, w io.Writer, depth int) {

	srcName := config.SrcName
	cursorName := config.CursorName
	hasConflict := tree.HasConflict(config)

	if config.CaseSensitive || hasConflict {
		this.GenerateSwitch(config, w, "switch (*(%s++))", cursorName)
	} else {
		this.GenerateSwitch(config, w, "switch (*(%s++) | 0x20)", cursorName)
	}

	for _, v := range tree.Branches {
		if v.Value[0] == 0 {
			continue
		}
		this.Fprintfln(w, "case %s:", getCharPrint(v.Value[0]))
		if chars.IsAlpha(v.Value[0]) && !config.CaseSensitive && hasConflict {
			this.Fprintfln(w, "case '%c':", chars.ToUpper(v.Value[0]))
		}
		this.EnterIndent(config.IndentOfBlock)
		this.GenerateActionCode(config, v.Next, w, depth+1)

		//this.GenerateBlockBegin(config, w, "")
		this.Fprintfln(w, "*%s = %s;", srcName, cursorName)
		this.Fprintfln(w, "return %s;", config.UnknownIndexName)
		//this.GenerateBlockEnd(config, w)

		this.Exit()
	}

	this.GenerateSwitchEnd(config, w)
}

func (this *KeyCmpGeneratorForCpp) GenerateActionCodeForSingleBranch(config *key_gen.Config,
	tree *key_cmp_gen.TrieTree, w io.Writer, depth int) {

	branch := tree.FirstNonFinalBranch()

	if len(branch.Value) == 1 {
		this.GenerateActionCodeForSingleBranchWithOneChar(config, branch, w, depth)
	} else {
		this.GenerateActionCodeForSingleBranchWithOneString(config, branch, w, depth)
	}
}

func (this *KeyCmpGeneratorForCpp) GenerateActionCodeForSingleBranchWithOneChar(config *key_gen.Config,
	branch *key_cmp_gen.Branch, w io.Writer, depth int) {

	cursorName := config.CursorName

	ch := branch.Value[0]
	chStr := getCharPrint(ch)

	if chars.IsAlpha(ch) && !config.CaseSensitive {
		this.Fprintf(w, "if ((%s < end) && ((*(%s++) | 0x20) == %s)", cursorName, cursorName, chStr)
	} else {
		this.Fprintf(w, "if ((%s < end) && (*(%s++) == %s)", cursorName, cursorName, chStr)
	}
	this.generateLeftBrace(config, w, config.IndentOfIf)
	this.GenerateActionCode(config, branch.Next, w, depth+1)
	this.generateRightBrace(config, w)
}

func (this *KeyCmpGeneratorForCpp) GenerateActionCodeForSingleBranchWithOneString(config *key_gen.Config,
	branch *key_cmp_gen.Branch, w io.Writer, depth int) {

	srcName := config.SrcName
	cursorName := config.CursorName

	this.Fprintf(w, "if ((%s + %d) >= end)", cursorName, len(branch.Value)-1)
	this.generateLeftBrace(config, w, config.IndentOfBlock)
	this.Fprintfln(w, "*%s = %s;", srcName, cursorName)
	this.Fprintfln(w, "return %s;", config.UnknownIndexName)
	this.generateRightBrace(config, w)

	if chars.IsAlpha(branch.Value[0]) && !config.CaseSensitive {
		this.Fprintf(w, "if (((*(%s++) | 0x20) == %s)", cursorName, getCharPrint(branch.Value[0]))
	} else {
		this.Fprintf(w, "if ((*(%s++) == %s)", cursorName, getCharPrint(branch.Value[0]))
	}
	this.EnterIndent(config.IndentOfBlock)
	for i := 1; i < len(branch.Value); i++ {
		fmt.Fprintln(w)
		if chars.IsAlpha(branch.Value[i]) && !config.CaseSensitive {
			this.Fprintf(w, "&& ((*(%s++) | 0x20) == %s)", cursorName, getCharPrint(branch.Value[i]))
		} else {
			this.Fprintf(w, "&& (*(%s++) == %s)", cursorName, getCharPrint(branch.Value[i]))
		}
	}

	fmt.Fprintf(w, ")")
	this.Exit()
	this.generateLeftBrace(config, w, config.IndentOfIf)
	this.GenerateActionCode(config, branch.Next, w, depth+1)
	this.generateRightBrace(config, w)
}

func getCharPrint(ch byte) string {
	if chars.IsPrintAscii(ch) {
		return fmt.Sprintf("'%c'", ch)
	}
	return fmt.Sprintf("0x%02x", ch)
}

func getIndexTypeName(config *key_gen.Config) string {
	if len(config.IndexTypeName) > 0 {
		return config.IndexTypeName
	}

	return getTypeNameBySize(config.IndexTypeSize)
}

func getSrcTypeName(config *key_gen.Config) string {
	if len(config.SrcTypeName) > 0 {
		return config.SrcTypeName
	}

	return "char const*"
}

func (this *KeyCmpGeneratorForCpp) GenerateSwitch(config *key_gen.Config, w io.Writer, format string, args ...interface{}) {
	this.Fprintf(w, format, args...)
	this.generateLeftBrace(config, w, config.IndentOfSwitch)
}

func (this *KeyCmpGeneratorForCpp) GenerateSwitchEnd(config *key_gen.Config, w io.Writer) {
	this.generateRightBrace(config, w)
}

func (this *KeyCmpGeneratorForCpp) GenerateBlockBegin(config *key_gen.Config, w io.Writer, format string, args ...interface{}) {
	this.Fprintf(w, format, args...)
	this.generateBlockLeftBrace(config, w, config.IndentOfBlock)
}

func (this *KeyCmpGeneratorForCpp) GenerateBlockEnd(config *key_gen.Config, w io.Writer) {
	this.generateRightBrace(config, w)
}

func (this *KeyCmpGeneratorForCpp) generateRightBrace(config *key_gen.Config, w io.Writer) {
	this.Exit()
	this.Fprintln(w, "}")
}

func (this *KeyCmpGeneratorForCpp) generateLeftBrace(config *key_gen.Config, w io.Writer, indent int) {
	if config.BraceAtNextLine {
		fmt.Fprintln(w)
		this.Fprintln(w, "{")
	} else {
		fmt.Fprintln(w, " {")
	}
	this.EnterIndent(indent)
}

func (this *KeyCmpGeneratorForCpp) generateBlockLeftBrace(config *key_gen.Config, w io.Writer, indent int) {
	fmt.Fprintln(w)
	this.Fprintln(w, "{")
	this.EnterIndent(indent)
}

func getTypeNameBySize(typeSize int) string {
	switch typeSize {
	case 1:
		return "unsigned char"
	case 2:
		return "unsigned short"
	case 8:
		return "unsigned long"
	default:
		return "unsigned int"
	}
}
