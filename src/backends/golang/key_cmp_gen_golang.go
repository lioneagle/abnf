package golang

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/lioneagle/abnf/src/basic"
	"github.com/lioneagle/abnf/src/gen/key_gen"
	"github.com/lioneagle/abnf/src/gen/key_gen/key_cmp_gen"
	"github.com/lioneagle/abnf/src/keys"

	"github.com/lioneagle/goutil/src/chars"
	"github.com/lioneagle/goutil/src/logger"
	"github.com/lioneagle/goutil/src/times"
)

type KeyCmpGeneratorForGolang struct {
	chars.Indent
}

func NewKeyCmpGeneratorForGolang() *KeyCmpGeneratorForGolang {
	ret := &KeyCmpGeneratorForGolang{}
	ret.Indent.Init(0, 4)
	ret.Indent.UseTab = true
	ret.Indent.SetReturnString("\n")
	return ret
}

func (this *KeyCmpGeneratorForGolang) GenerateFile(config *key_gen.Config,
	keys *keys.Keys, filename, path string) {

	timeStat := times.NewTimeStat()
	defer func() {
		timeStat.Stop()
		if config.PrintTimeUsed {
			timeStat.Fprint(config.OutputFile, "tokencmp generate files")
		}
	}()

	this.generateFile(config, keys, filename, path)
}

func (this *KeyCmpGeneratorForGolang) buildTrieTree(config *key_gen.Config, keys *keys.Keys) *key_cmp_gen.TrieTree {
	timeStat := times.NewTimeStat()
	defer func() {
		timeStat.Stop()
		if config.PrintTimeUsed {
			timeStat.Fprint(config.OutputFile, "tokencmp build trie tree")
		}
	}()

	return key_cmp_gen.BuildTrieTreeFromKeys(config, keys)
}

func (this *KeyCmpGeneratorForGolang) generateFile(config *key_gen.Config,
	keys *keys.Keys, filename, path string) {

	timeStat := times.NewTimeStat()
	defer func() {
		timeStat.Stop()
		if config.PrintTimeUsed {
			timeStat.Fprint(config.OutputFile, "tokencmp generate go file")
		}
	}()

	abs_filename := filepath.FromSlash(path + "/" + filename + ".go")
	file, err := os.Create(abs_filename)
	if err != nil {
		logger.Error("cannot open file %s", abs_filename)
		return
	}
	defer file.Close()

	tree := this.buildTrieTree(config, keys)

	this.Fprintfln(file, "package %s", config.PackageName)
	this.Fprintln(file)

	this.Fprintln(file, "/*---------------- index definition ----------------*/")
	this.GenerateIndex(config, keys, file)
	this.Fprintln(file)

	this.GenerateActionDefinition(config, tree, file)
}

func (this *KeyCmpGeneratorForGolang) GenerateIndex(config *key_gen.Config,
	keys *keys.Keys, w io.Writer) {

	MaxIndexName := keys.GetMaxIndexNameLen()
	indexTypeName := getIndexTypeName(config)

	this.Fprintln(w, "const (")
	this.Enter()

	for _, v := range keys.Indices {
		this.Fprintf(w, "%s", v.Name)
		basic.PrintIndent(w, MaxIndexName-len(v.Name))
		fmt.Fprintf(w, " %s = %d", indexTypeName, v.Value)
		this.PrintReturn(w)
	}

	this.Exit()
	this.Fprintln(w, ")")
}

func (this *KeyCmpGeneratorForGolang) GenerateActionDeclaration(config *key_gen.Config, w io.Writer) {
	indexType := getIndexTypeName(config)
	srcType := getSrcTypeName(config)

	this.Fprintf(w, "func %s(%s %s) %s", config.ActionName, config.SrcName, srcType, indexType)
}

func (this *KeyCmpGeneratorForGolang) GenerateActionDefinition(config *key_gen.Config,
	tree *key_cmp_gen.TrieTree, w io.Writer) {

	srcName := config.SrcName

	this.GenerateActionDeclaration(config, w)

	this.generateLeftBrace(config, w, config.IndentOfBlock)

	this.Fprintfln(w, "%s := 0", config.CursorName)
	this.Fprintfln(w, "len1 := len(%s)", srcName)
	this.Fprintln(w)

	this.Fprintf(w, "if %s >= len1", config.CursorName)
	this.generateLeftBrace(config, w, config.IndentOfBlock)
	this.Fprintfln(w, "return %s", config.UnknownIndexName)
	this.generateRightBrace(config, w)
	this.Fprintln(w)

	this.GenerateActionCode(config, tree, w, 0)

	this.Fprintln(w)
	this.Fprintfln(w, "return %s", config.UnknownIndexName)

	this.GenerateBlockEnd(config, w)
}

func (this *KeyCmpGeneratorForGolang) GenerateActionCode(config *key_gen.Config,
	tree *key_cmp_gen.TrieTree, w io.Writer, depth int) {

	srcName := config.SrcName
	seperatorName := config.SeperatorCharsetName
	cursorName := config.CursorName

	branch := tree.FindFinal(config)

	if branch != nil {
		if config.SeperatorEnabled {
			this.Fprintf(w, "if (%s < len1) && %s(%s[%s])", cursorName, seperatorName, srcName, cursorName)
		} else {
			this.Fprintf(w, "if %s >= len1", cursorName)
		}
		this.generateLeftBrace(config, w, config.IndentOfBlock)

		//this.Fprintfln(w, "*%s = %s;", srcName, cursorName)
		this.Fprintfln(w, "return %s", branch.Next.Key.Index.Name)
		this.generateRightBrace(config, w)

		if depth == 0 && len(tree.Branches) == 1 {
			this.GenerateBlockBegin(config, w, "")
			//this.Fprintfln(w, "*%s = %s;", srcName, cursorName)
			this.Fprintfln(w, "return %s", config.UnknownIndexName)
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

func (this *KeyCmpGeneratorForGolang) GenerateActionCodeForMultiBranch(config *key_gen.Config,
	tree *key_cmp_gen.TrieTree, w io.Writer, depth int) {

	srcName := config.SrcName
	cursorName := config.CursorName
	hasConflict := tree.HasConflict(config)

	if config.CaseSensitive || hasConflict {
		this.GenerateSwitch(config, w, "switch %s[%s]", srcName, cursorName)
	} else {
		this.GenerateSwitch(config, w, "switch %s[%s] | 0x20", srcName, cursorName)
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
		this.Fprintfln(w, "%s++", cursorName)
		this.GenerateActionCode(config, v.Next, w, depth+1)

		//this.GenerateBlockBegin(config, w, "")
		//this.Fprintfln(w, "*%s = %s;", srcName, cursorName)
		this.Fprintfln(w, "return %s", config.UnknownIndexName)
		//this.GenerateBlockEnd(config, w)

		this.Exit()
	}

	this.GenerateSwitchEnd(config, w)
}

func (this *KeyCmpGeneratorForGolang) GenerateActionCodeForSingleBranch(config *key_gen.Config,
	tree *key_cmp_gen.TrieTree, w io.Writer, depth int) {

	branch := tree.FirstNonFinalBranch()

	if len(branch.Value) == 1 {
		this.GenerateActionCodeForSingleBranchWithOneChar(config, branch, w, depth)
	} else {
		this.GenerateActionCodeForSingleBranchWithOneString(config, branch, w, depth)
	}
}

func (this *KeyCmpGeneratorForGolang) GenerateActionCodeForSingleBranchWithOneChar(config *key_gen.Config,
	branch *key_cmp_gen.Branch, w io.Writer, depth int) {

	srcName := config.SrcName
	cursorName := config.CursorName

	ch := branch.Value[0]
	chStr := getCharPrint(ch)

	if chars.IsAlpha(ch) && !config.CaseSensitive {
		this.Fprintf(w, "if (%s < len1) && ((%s[%s] | 0x20) == %s)", cursorName, srcName, cursorName, chStr)
	} else {
		this.Fprintf(w, "if (%s < len1) && (%s[%s] == %s)", cursorName, srcName, cursorName, chStr)
	}
	this.generateLeftBrace(config, w, config.IndentOfIf)
	this.Fprintfln(w, "%s++", cursorName)
	this.GenerateActionCode(config, branch.Next, w, depth+1)
	this.generateRightBrace(config, w)
}

func (this *KeyCmpGeneratorForGolang) GenerateActionCodeForSingleBranchWithOneString(config *key_gen.Config,
	branch *key_cmp_gen.Branch, w io.Writer, depth int) {

	srcName := config.SrcName
	cursorName := config.CursorName

	this.Fprintf(w, "if (%s + %d) >= len1", cursorName, len(branch.Value)-1)
	this.generateLeftBrace(config, w, config.IndentOfBlock)
	//this.Fprintfln(w, "*%s = %s;", srcName, cursorName)
	this.Fprintfln(w, "return %s", config.UnknownIndexName)
	this.generateRightBrace(config, w)

	if chars.IsAlpha(branch.Value[0]) && !config.CaseSensitive {
		this.Fprintf(w, "if ((%s[%s] | 0x20) == %s)", srcName, cursorName, getCharPrint(branch.Value[0]))
	} else {
		this.Fprintf(w, "if (%s[%s] == %s)", srcName, cursorName, getCharPrint(branch.Value[0]))
	}
	this.EnterIndent(config.IndentOfBlock)
	for i := 1; i < len(branch.Value); i++ {
		fmt.Fprintln(w, " &&")
		if chars.IsAlpha(branch.Value[i]) && !config.CaseSensitive {
			this.Fprintf(w, "((%s[%s+%d] | 0x20) == %s)", srcName, cursorName, i, getCharPrint(branch.Value[i]))
		} else {
			this.Fprintf(w, "(%s[%s+%d] == %s)", srcName, cursorName, i, getCharPrint(branch.Value[i]))
		}
	}

	this.Exit()
	this.generateLeftBrace(config, w, config.IndentOfIf)
	this.Fprintfln(w, "%s += %d", cursorName, len(branch.Value))
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

	return "[]byte"
}

func (this *KeyCmpGeneratorForGolang) GenerateSwitch(config *key_gen.Config, w io.Writer, format string, args ...interface{}) {
	this.Fprintf(w, format, args...)
	this.generateLeftBrace(config, w, config.IndentOfSwitch)
	this.Exit()
}

func (this *KeyCmpGeneratorForGolang) GenerateSwitchEnd(config *key_gen.Config, w io.Writer) {
	this.Fprintln(w, "}")
}

func (this *KeyCmpGeneratorForGolang) GenerateBlockBegin(config *key_gen.Config, w io.Writer, format string, args ...interface{}) {
	this.Fprintf(w, format, args...)
	this.generateBlockLeftBrace(config, w, config.IndentOfBlock)
}

func (this *KeyCmpGeneratorForGolang) GenerateBlockEnd(config *key_gen.Config, w io.Writer) {
	this.generateRightBrace(config, w)
}

func (this *KeyCmpGeneratorForGolang) generateRightBrace(config *key_gen.Config, w io.Writer) {
	this.Exit()
	this.Fprintln(w, "}")
}

func (this *KeyCmpGeneratorForGolang) generateLeftBrace(config *key_gen.Config, w io.Writer, indent int) {
	if config.BraceAtNextLine {
		fmt.Fprintln(w)
		this.Fprintln(w, "{")
	} else {
		fmt.Fprintln(w, " {")
	}
	this.EnterIndent(indent)
}

func (this *KeyCmpGeneratorForGolang) generateBlockLeftBrace(config *key_gen.Config, w io.Writer, indent int) {
	fmt.Fprintln(w)
	this.Fprintln(w, "{")
	this.EnterIndent(indent)
}

func getTypeNameBySize(typeSize int) string {
	switch typeSize {
	case 1:
		return "byte"
	case 2:
		return "uint16"
	case 8:
		return "uint64"
	default:
		return "uint32"
	}
}
