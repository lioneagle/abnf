package peg_gen

import (
	"github.com/lioneagle/abnf/src/gen"
)

type Config struct {
	gen.ConfigBase

	PadTypeName string
}

func NewConfig() *Config {
	ret := &Config{}
	ret.Init()
	return ret
}

func (this *Config) Init() {
	this.ConfigBase.Init()
	this.PadTypeName = "byte"
}
