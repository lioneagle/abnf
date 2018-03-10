package keys

import (
	//"fmt"
	"sort"

	"github.com/lioneagle/goutil/src/logger"
)

type Key struct {
	Name  string
	Index Index
	//IndexName     string
	//IndexValue    int
	LineNo        int
	CaseSensitive bool
}

type Index struct {
	Name  string
	Value int
}

type Keys struct {
	Data     []*Key
	IndexMap map[string]int
	Indices  []*Index
}

func NewKeys() *Keys {
	return &Keys{IndexMap: make(map[string]int)}
}

func (this *Keys) Add(key *Key) {
	this.Data = append(this.Data, key)
	this.AddIndex(&key.Index)
}

func (this *Keys) AddIndex(index *Index) {
	if len(index.Name) <= 0 {
		return
	}

	val, ok := this.IndexMap[index.Name]
	if !ok {
		this.IndexMap[index.Name] = index.Value
		this.Indices = append(this.Indices, index)
		sort.Slice(this.Indices, func(i, j int) bool { return this.Indices[i].Value < this.Indices[j].Value })
		return
	}

	if val != index.Value {
		logger.Error("different value for \"%s\", current = %d, old = %d", index.Name, index.Value, val)
	}

}

func (this *Keys) GetMaxMinNameLen() (max, min int) {
	max = 0
	min = 0
	for _, v := range this.Data {
		len1 := len(v.Name)
		if len1 > max {
			max = len1
		}
		if len1 < min {
			min = len1
		}
	}
	return max, min
}

func (this *Keys) GetMaxIndexNameLen() (max int) {
	max = 0
	for _, v := range this.Data {
		len1 := len(v.Index.Name)
		if len1 > max {
			max = len1
		}
	}
	return max
}
