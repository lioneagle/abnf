package key_gperf_gen

import (
	"github.com/lioneagle/abnf/src/keys"

	"github.com/lioneagle/goutil/src/logger"
)

const (
	MAX_POSITIONS = 256
	LAST_CHAR     = -1
)

// Positions stores a set of positions, used to access a keyword
type Positions struct {
	Data []int
}

func (this *Positions) GetFromKeys(keys *keys.Keys) bool {
	return false
}

func (this *Positions) GetMandatoryPositionsFromKeys(keys *keys.Keys) bool {
	return false
}

func (this *Positions) Contains(pos int) bool {
	for _, v := range this.Data {
		if v == pos {
			return true
		}
	}
	return false
}

func (this *Positions) Add(pos int) bool {
	if this.Contains(pos) {
		logger.Error("Positions.Add: pos(=%d) duplicated", pos)
		return false
	}

	this.Data = append(this.Data, pos)
	return false
}
