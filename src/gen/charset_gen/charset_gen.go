package charset_gen

import (
	"fmt"
	"io"
	"strings"

	"github.com/lioneagle/abnf/src/charset"
)

type CharsetInfo struct {
	Name    string
	Charset *charset.Charset

	MaskValue uint64
	VarIndex  int
}

func NewCharsetInfo(name string) *CharsetInfo {
	return &CharsetInfo{Name: name}
}

func (this *CharsetInfo) GetMaskName(config *Config) string {
	if config.MaskPrefix == "" {
		return fmt.Sprintf("MASK_%s", strings.Replace(strings.ToUpper(this.Name), "-", "_", -1))
	}
	return fmt.Sprintf("%s_MASK_%s", config.MaskPrefix, strings.Replace(strings.ToUpper(this.Name), "-", "_", -1))
}

func (this *CharsetInfo) GetActionName(config *Config) string {
	act := "Is"
	if config.ActionFirstLower {
		act = "is"
	}
	if config.ActionPrefix == "" {
		return fmt.Sprintf("%s%s", act, strings.Replace(strings.Title(this.Name), "-", "", -1))
	}
	return fmt.Sprintf("%s_%s%s", config.ActionPrefix, act, strings.Replace(strings.Title(this.Name), "-", "", -1))
}

type Var struct {
	Data [256]uint64
}

func (this *Var) SetCharset(c *charset.Charset, mask uint64) {
	for i := int32(0); i < 256; i++ {
		if c.Contains(i) {
			this.Data[i] |= mask
		}
	}
}

type CharsetTable struct {
	Name             string
	Charsets         []*CharsetInfo
	Vars             []Var
	MaskNameMaxLen   int
	ActionNameMaxLen int
}

func NewCharsetTable() *CharsetTable {
	return &CharsetTable{Name: "charset"}
}

func (this *CharsetTable) Add(val *CharsetInfo) {
	this.Charsets = append(this.Charsets, val)
}

func (this *CharsetTable) Calc(config *Config) {
	if len(this.Charsets) == 0 {
		return
	}

	this.calcNameMaxLen(config)
	this.calcVar(config)

}

func (this *CharsetTable) calcNameMaxLen(config *Config) {
	this.MaskNameMaxLen = 0
	this.ActionNameMaxLen = 0
	for _, v := range this.Charsets {
		maskNameLen := len(v.GetMaskName(config))
		actionNameLen := len(v.GetActionName(config))

		if maskNameLen > this.MaskNameMaxLen {
			this.MaskNameMaxLen = maskNameLen
		}

		if actionNameLen > this.ActionNameMaxLen {
			this.ActionNameMaxLen = actionNameLen
		}
	}
}

func (this *CharsetTable) calcVar(config *Config) {
	if config.UseBit {
		this.calcVarUseBit(config)
	} else {
		this.calcVarUseByte(config)
	}
}

func (this *CharsetTable) calcVarUseBit(config *Config) {
	typeBit := config.VarTypeSize * 8
	charsetNum := len(this.Charsets)
	varNum := charsetNum / typeBit
	if (charsetNum % typeBit) != 0 {
		varNum++
	}

	this.Vars = make([]Var, varNum)

	for i, v := range this.Charsets {
		v.MaskValue = 1 << uint64(i%typeBit)
		v.VarIndex = i / typeBit
		this.Vars[v.VarIndex].SetCharset(v.Charset, v.MaskValue)
	}

}

func (this *CharsetTable) calcVarUseByte(config *Config) {
	this.Vars = make([]Var, len(this.Charsets))
	for i, v := range this.Charsets {
		v.VarIndex = i
		this.Vars[i].SetCharset(v.Charset, 1)
	}
}

type CharsetTableGenerator interface {
	GenerateFile(config *Config, charsets *CharsetTable, filename, path string)
	GenerateMask(config *Config, charsets *CharsetTable, w io.Writer)
	GenerateAction(config *Config, charsets *CharsetTable, w io.Writer)
	GenerateVarDeclaration(config *Config, charsets *CharsetTable, w io.Writer)
	GenerateVarDefinition(config *Config, charsets *CharsetTable, w io.Writer)
}
