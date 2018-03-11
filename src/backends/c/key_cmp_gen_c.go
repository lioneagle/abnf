package c

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	//"strconv"
	"strings"

	"github.com/lioneagle/abnf/src/basic"
	"github.com/lioneagle/abnf/src/gen/key_gen"
	"github.com/lioneagle/abnf/src/gen/key_gen/key_cmp_gen"
	"github.com/lioneagle/abnf/src/keys"

	"github.com/lioneagle/goutil/src/chars"
	"github.com/lioneagle/goutil/src/logger"
	"github.com/lioneagle/goutil/src/times"
)

type KeyCmpGeneratorForC struct {
	chars.Indent
}

func NewKeyCmpGeneratorForC() *KeyCmpGeneratorForC {
	ret := &KeyCmpGeneratorForC{}
	ret.Indent.Init(0, 4)
	return ret
}

func (this *KeyCmpGeneratorForC) GenerateFile(config *key_gen.Config,
	keys *keys.Keys, filename, path string) {

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

func (this *KeyCmpGeneratorForC) buildTrieTree(config *key_gen.Config, keys *keys.Keys) *key_cmp_gen.TrieTree {
	timeStat := times.NewTimeStat()
	defer func() {
		timeStat.Stop()
		if config.PrintTimeUsed {
			timeStat.Fprint(config.OutputFile, "tokencmp build trie tree")
		}
	}()

	return key_cmp_gen.BuildTrieTreeFromKeys(config, keys)
}

func (this *KeyCmpGeneratorForC) generateHFile(config *key_gen.Config,
	keys *keys.Keys, filename, path string) {

	timeStat := times.NewTimeStat()
	defer func() {
		timeStat.Stop()
		if config.PrintTimeUsed {
			timeStat.Fprint(config.OutputFile, "tokencmp generate h file")
		}
	}()

	abs_filename := filepath.FromSlash(path + "/" + filename + ".h")
	file, err := os.Create(abs_filename)
	if err != nil {
		logger.Error("cannot open file %s", abs_filename)
		return
	}
	defer file.Close()

	name := strings.ToUpper(filename)

	this.Fprintfln(file, "#ifndef %s_H", name)
	this.Fprintfln(file, "#define %s_H", name)
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

	this.Fprintfln(file, "#endif /* %s_H */", name)
}

func (this *KeyCmpGeneratorForC) generateCFile(config *key_gen.Config,
	tree *key_cmp_gen.TrieTree, filename, path string) {

	timeStat := times.NewTimeStat()
	defer func() {
		timeStat.Stop()
		if config.PrintTimeUsed {
			timeStat.Fprint(config.OutputFile, "tokencmp generate c file")
		}
	}()

	abs_filename := filepath.FromSlash(path + "/" + filename + ".c")
	file, err := os.Create(abs_filename)
	if err != nil {
		logger.Error("cannot open file %s", abs_filename)
		return
	}
	defer file.Close()

	this.Fprintfln(file, "#include \"%s\"", filename+".h")
	this.Fprintln(file)
	this.GenerateActionDefinition(config, tree, file)
}

func (this *KeyCmpGeneratorForC) GenerateIndex(config *key_gen.Config,
	keys *keys.Keys, w io.Writer) {

	MaxIndexName := keys.GetMaxIndexNameLen()

	format := fmt.Sprintf("((%s)(%%d))", getIndexTypeName(config))
	for _, v := range keys.Indices {
		this.Fprintf(w, "#define %s", v.Name)
		basic.PrintIndent(w, MaxIndexName+4-len(v.Name))
		this.Fprintfln(w, format, v.Value)
	}
}

func (this *KeyCmpGeneratorForC) GenerateActionDeclaration(config *key_gen.Config, w io.Writer) {
	indexType := getIndexTypeName(config)
	srcType := getSrcTypeName(config)

	this.Fprintf(w, "%s %s(%s* %s, %s end)", indexType, config.ActionName, srcType, config.SrcName, srcType)
}

func (this *KeyCmpGeneratorForC) GenerateActionDefinition(config *key_gen.Config,
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

	fmt.Println("****************************************")
	tree.Print(0)
	fmt.Println("****************************************")

	this.GenerateActionCode(config, tree, w, 0)

	this.Fprintln(w)
	this.Fprintfln(w, "return %s;", config.UnknownIndexName)

	this.GenerateBlockEnd(config, w)
}

func (this *KeyCmpGeneratorForC) GenerateActionCode(config *key_gen.Config,
	tree *key_cmp_gen.TrieTree, w io.Writer, depth int) {

	srcName := config.SrcName
	seperatorName := config.SeperatorCharsetName
	cursorName := config.CursorName

	branch := tree.FindFinal(config)

	tree.PrintBranches(depth)

	if branch != nil {
		fmt.Println("there1")
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
			fmt.Println("there1.1")
			this.GenerateBlockBegin(config, w, "")
			this.Fprintfln(w, "*%s = %s;", srcName, cursorName)
			this.Fprintfln(w, "return %s;", config.UnknownIndexName)
			this.GenerateBlockEnd(config, w)
			return
		}
	}

	/*if (len(tree.Branches) == 1 && tree.Branches[0].Value[0] != 0) ||
	(len(tree.Branches) == 2 && tree.Branches[0].Value[0] == 0) {

	if len(tree.Branches) == 1 && tree.Branches[0].Value[0] != 0 {
		branch = tree.Branches[0]
	} else {
		branch = tree.Branches[1]
	}*/
	if tree.NonFinalBranchNum() == 1 {
		fmt.Println("there2")
		branch = tree.FirstNonFinalBranch()

		fmt.Println("branch.Value =", string(branch.Value))

		if len(branch.Value) == 1 {
			fmt.Println("there3")
			ch := branch.Value[0]
			chStr := getCharPrint(ch)

			if chars.IsAlpha(ch) && !config.CaseSensitive {
				this.Fprintf(w, "if ((%s < end) && ((*(%s++) | 0x20) == %s)", srcName, srcName, chStr)
			} else {
				this.Fprintf(w, "if ((%s < end) && (*(%s++) == %s)", cursorName, cursorName, chStr)
			}
			this.generateLeftBrace(config, w, config.IndentOfIf)
			this.GenerateActionCode(config, branch.Next, w, depth+1)
			this.generateRightBrace(config, w)
		} else {
			fmt.Println("there4")
			this.Fprintf(w, "if ((%s+%d) >= end)", cursorName, len(branch.Value)-1)
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
	} else if tree.NonFinalBranchNum() > 1 {
		fmt.Println("there5")

		hasConflict := tree.HasConflict(config)

		//fmt.Println("len(tree.Branches) =", len(tree.Branches))

		if config.CaseSensitive || hasConflict {
			this.GenerateSwitch(config, w, "switch (*(%s++))", cursorName)
		} else {
			this.GenerateSwitch(config, w, "switch (*(%s++) | 0x20)", cursorName)
		}
		for _, v := range tree.Branches {
			if v.Value[0] == 0 {
				continue
			}
			fmt.Printf("v.Value = %s, v.Next = %p\n", string(v.Value), v.Next)
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
}

/*func (this *KeyCmpGeneratorForC) GenerateStringCode(config *key_gen.Config,
	branch *key_cmp_gen.Branch, w io.Writer, depth int, value []byte) {

	if len(branch.Value) == 1 {
		ch := branch.Value[0]
		chStr := getCharPrint(ch)

		if chars.IsAlpha(ch) && !config.CaseSensitive {
			this.Fprintf(w, "if ((%s < end) && ((*(%s++) | 0x20) == %s)", srcName, srcName, chStr)
		} else {
			this.Fprintf(w, "if ((%s < end) && (*(%s++) == %s)", cursorName, cursorName, chStr)
		}
		this.generateLeftBrace(config, w, config.IndentOfIf)
		this.GenerateActionCode(config, branch.Next, w, depth+1)
		this.generateRightBrace(config, w)
	} else {
		this.Fprintf(w, "if ((%s+%d) >= end)", cursorName, len(branch.Value)-1)
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
			this.Fprintln(w)
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
}*/

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

func (this *KeyCmpGeneratorForC) GenerateSwitch(config *key_gen.Config, w io.Writer, format string, args ...interface{}) {
	this.Fprintf(w, format, args...)
	this.generateLeftBrace(config, w, config.IndentOfSwitch)
}

func (this *KeyCmpGeneratorForC) GenerateSwitchEnd(config *key_gen.Config, w io.Writer) {
	this.generateRightBrace(config, w)
}

func (this *KeyCmpGeneratorForC) GenerateBlockBegin(config *key_gen.Config, w io.Writer, format string, args ...interface{}) {
	this.Fprintf(w, format, args...)
	this.generateBlockLeftBrace(config, w, config.IndentOfBlock)
}

func (this *KeyCmpGeneratorForC) GenerateBlockEnd(config *key_gen.Config, w io.Writer) {
	this.generateRightBrace(config, w)
}

func (this *KeyCmpGeneratorForC) generateRightBrace(config *key_gen.Config, w io.Writer) {
	this.Exit()
	this.Fprintln(w, "}")
}

func (this *KeyCmpGeneratorForC) generateLeftBrace(config *key_gen.Config, w io.Writer, indent int) {
	if config.BraceAtNextLine {
		fmt.Fprintln(w)
		this.Fprintln(w, "{")
	} else {
		fmt.Fprintln(w, " {")
	}
	this.EnterIndent(indent)
}

func (this *KeyCmpGeneratorForC) generateBlockLeftBrace(config *key_gen.Config, w io.Writer, indent int) {
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
