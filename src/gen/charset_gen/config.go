package charset_gen

import (
	"strings"
)

type Config struct {
	MaskPrefix       string
	ActionPrefix     string
	ActionFirstLower bool
	UseBit           bool
	VarName          string
	VarTypeName      string
	VarTypeSize      int
	PackageName      string
}

func NewConfig() *Config {
	return &Config{UseBit: true, VarName: "charsets", VarTypeSize: 4}
}

func (this *Config) SetMaskPrefix(prefix string) {
	this.MaskPrefix = strings.ToUpper(prefix)
}

func (this *Config) SetActionPrefix(prefix string) {
	this.ActionPrefix = prefix
}

func (this *Config) SetVarTypeSize(val int) {
	switch val {
	case 1, 2, 4, 8:
		this.VarTypeSize = val
	default:
		this.VarTypeSize = 4
	}
}

func (this *Config) SetVarName(val string) {
	if len(val) == 0 {
		this.VarName = "charsets"
	} else {
		this.VarName = val
	}
}
