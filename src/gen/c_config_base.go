package gen

type CConfigBase struct {
	ConfigBase
}

func NewCConfigBase() *CConfigBase {
	ret := &CConfigBase{}
	ret.Init()
	return ret
}

func (this *CConfigBase) Init() {
	this.ConfigBase.Init()
}
