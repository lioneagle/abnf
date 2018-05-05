package key_gen

import (
	"io"
	"os"
)

type Config struct {
	CaseSensitive        bool
	SeperatorEnabled     bool
	CharsetEnabled       bool
	BraceAtNextLine      bool
	PrintTimeUsed        bool
	BuildSimpleTree      bool
	UseTabIndent         bool
	GenVersion           bool
	SeperatorCharsetName string
	CharsetName          string
	ActionName           string
	IndexTypeName        string
	IndexTypeSize        int
	SrcTypeName          string
	SrcName              string
	CursorName           string
	CursorTypeName       string
	CursorTypeSize       int
	PackageName          string
	UnknownIndexName     string
	UnknownIndexValue    int
	MaxIndexName         string
	IndentOfIf           int
	IndentOfSwitch       int
	IndentOfBlock        int
	OutputFile           io.Writer
	ErrorFile            io.Writer
	DebugFile            io.Writer
}

func NewConfig() *Config {
	ret := &Config{}
	ret.CaseSensitive = true
	ret.ActionName = "Lookup"
	ret.SrcName = "src"
	ret.CursorName = "p"
	ret.CursorTypeName = "int"
	ret.CursorTypeSize = 4
	ret.UnknownIndexName = "UNKNOWN"
	ret.IndexTypeSize = 4
	ret.BraceAtNextLine = true
	ret.PrintTimeUsed = true
	ret.OutputFile = os.Stdout
	ret.ErrorFile = os.Stderr
	//ret.IndentOfIf = 2
	ret.IndentOfIf = 4
	ret.IndentOfSwitch = 4
	ret.IndentOfBlock = 4
	return ret
}

func (this *Config) SetCaseSensitive() {
	this.CaseSensitive = true
}

func (this *Config) SetCaseInsensitive() {
	this.CaseSensitive = false
}
