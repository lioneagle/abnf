package gen

import (
	"fmt"
	"io"
	"time"
)

const (
	VERSION = "v0.1.0"
)

type Var struct {
	TypeName  string
	Name      string
	Comment   string
	InitValue string
}

func NewVar() *Var {
	return &Var{}
}

type ParamList struct {
	Params []*Var
}

func NewParamList() *ParamList {
	return &ParamList{}
}

func (this *ParamList) Append(param *Var) {
	this.Params = append(this.Params, param)
}

func (this *ParamList) GetMaxNameLen() int {
	max := 0
	for _, v := range this.Params {
		if len(v.Name) > max {
			max = len(v.Name)
		}
	}
	return max
}

func (this *ParamList) GetMaxTypeNameLen() int {
	max := 0
	for _, v := range this.Params {
		if len(v.TypeName) > max {
			max = len(v.TypeName)
		}
	}
	return max
}

func (this *ParamList) GetMaxValueLen() int {
	max := 0
	for _, v := range this.Params {
		if len(v.InitValue) > max {
			max = len(v.InitValue)
		}
	}
	return max
}

type Enum struct {
	Name  string
	Enums *ParamList
}

func NewEnum() *Enum {
	return &Enum{Enums: NewParamList()}
}

func (this *Enum) AppendEnum(enum *Var) {
	this.Enums.Append(enum)
}

type Struct struct {
	Name   string
	Fields *ParamList
}

func NewStruct() *Struct {
	return &Struct{Fields: NewParamList()}
}

func (this *Struct) AppendField(field *Var) {
	this.Fields.Append(field)
}

type Function struct {
	Name       string
	ReturnType string
	Params     *ParamList
	Comment    string
}

func NewFunction() *Function {
	return &Function{}
}

func (this *Function) AppendParam(param *Var) {
	this.Params.Append(param)
}

type Block struct {
	Code string
}

type Choice struct {
	Condition string
	Block     *Block
}

type SingleChoice struct {
	Condition  string
	BlockTrue  *Block
	BlockFalse *Block
}

type MultiChoice struct {
	Choices    []*Choice
	LastChoice *Choice
}

func GenerateVersion() string {
	return fmt.Sprintf("---------------- generated by abnf %s %s ----------------", VERSION, time.Now().Format(TIMESTAMP_FMT))
}

type GeneratorBase interface {
	GenerateFunctionDeclare(w io.Writer, f *Function)
	GenerateFunctionDefine(w io.Writer, f *Function)

	GenerateVarDeclare(w io.Writer, v *Var, typeInterval, nameInterval int)
	GenerateVarDefine(w io.Writer, v *Var, typeInterval, nameInterval int)

	GenerateVarEnum(w io.Writer, v *Var, nameInterval, valueInterval int)
	GenerateVarParam(w io.Writer, v *Var)

	GenerateParamList(w io.Writer, params *ParamList)

	GenerateMultiLineComment(w io.Writer, comment string)
	GenerateSingleLineComment(w io.Writer, comment string)
	GenerateSingleLineCommentWithoutIndent(w io.Writer, comment string)

	GenerateStructDefine(w io.Writer, s *Struct)
	GenerateEnumDefine(w io.Writer, e *Enum)

	GenerateSingleChoice(w io.Writer, c *SingleChoice)
	GenerateMultiChoice(w io.Writer, c *MultiChoice)

	GenerateBlockBegin(w io.Writer, indent int)
	GenerateBlockEnd(w io.Writer)

	GenerateForBegin(w io.Writer, format string, args ...interface{})
	GenerateForEnd(w io.Writer)
}
