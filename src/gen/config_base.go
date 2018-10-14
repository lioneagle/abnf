package gen

import (
	"io"
)

const TIMESTAMP_FMT string = "2006-01-02 15:04:05.000"

type Indents struct {
	If        int
	Switch    int
	Case      int
	While     int
	For       int
	DoWhile   int
	Block     int
	Struct    int
	FuncParam int
	Enum      int
	Union     int
	Assign    int
	Comment   int
}

func (this *Indents) Init() {
	this.If = 4
	this.Switch = 4
	this.Case = 4
	this.While = 4
	this.For = 4
	this.DoWhile = 4
	this.Block = 4
	this.Struct = 4
	this.FuncParam = 4
	this.Enum = 4
	this.Union = 4
	this.Assign = 1
	this.Comment = 1
}

type ConfigBase struct {
	Indent Indents

	GenVersion    bool
	PrintTimeUsed bool
	UseTabIndent  bool
	OutputPath    string
	DebugFileName string
	DebugFile     io.Writer

	BraceAtNextLine          bool
	VarUseSingleLineComment  bool
	ParamsInOneLine          bool
	MultiLineCommentDecorate bool
}

func NewConfigBase() *ConfigBase {
	ret := &ConfigBase{}
	ret.Init()
	return ret
}

func (this *ConfigBase) Init() {
	this.PrintTimeUsed = true
	this.OutputPath = "./"
	this.Indent.Init()
	this.BraceAtNextLine = true
	this.VarUseSingleLineComment = true
	this.ParamsInOneLine = true
}
