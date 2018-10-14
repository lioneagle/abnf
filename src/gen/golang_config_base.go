package gen

type GolangConfigBase struct {
	ConfigBase
	PackageName string
}

func NewGolangConfigBase() *GolangConfigBase {
	ret := &GolangConfigBase{}
	ret.Init()
	return ret
}

func (this *GolangConfigBase) Init() {
	this.ConfigBase.Init()
	this.Indent.Switch = 0
	this.PackageName = "abnf"
}
