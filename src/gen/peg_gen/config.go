package peg_gen

import (
	"io"
	"os"
)

type Config struct {
	BraceAtNextLine bool
	PrintTimeUsed   bool
	BuildSimpleTree bool
	UseTabIndent    bool
	GenVersion      bool
	PackageName     string
	IndentOfIf      int
	IndentOfSwitch  int
	IndentOfBlock   int
	PadTypeName     string
	OutputFile      io.Writer
	ErrorFile       io.Writer
	DebugFile       io.Writer
}

func NewConfig() *Config {
	ret := &Config{}
	ret.BraceAtNextLine = true
	ret.PrintTimeUsed = true
	ret.OutputFile = os.Stdout
	ret.ErrorFile = os.Stderr
	//ret.IndentOfIf = 2
	ret.IndentOfIf = 4
	ret.IndentOfSwitch = 4
	ret.IndentOfBlock = 4
	ret.PadTypeName = "byte"
	return ret
}
