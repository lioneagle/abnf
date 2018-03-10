package backends

import (
	"github.com/lioneagle/abnf/src/gen/key_gen"
	"github.com/lioneagle/abnf/src/keys"
)

func BuildKeysForTest(config *key_gen.Config) *keys.Keys {
	data := []struct {
		Name       string
		IndexName  string
		IndexValue int
	}{
		{"From", "ABNF_SIP_HDR_FROM", 1},
		{"f", "ABNF_SIP_HDR_FROM", 1},
		//{"To", "ABNF_SIP_HDR_TO", 2},
		//{"t", "ABNF_SIP_HDR_TO", 2},
		//{"Via", "ABNF_SIP_HDR_VIA", 3},
		//{"v", "ABNF_SIP_HDR_VIA", 3},
	}

	ret := keys.NewKeys()

	for _, v := range data {
		key := &keys.Key{Name: v.Name, Index: keys.Index{Name: v.IndexName, Value: v.IndexValue}}
		ret.Add(key)
	}

	if len(config.UnknownIndexName) > 0 {
		ret.AddIndex(&keys.Index{config.UnknownIndexName, config.UnknownIndexValue})
	}

	return ret
}

func BuildKeyGenConfigForTest() *key_gen.Config {
	config := key_gen.NewConfig()

	config.ActionName = "GetSipHeaderIndex"

	config.UnknownIndexName = "ABNF_SIP_HDR_UNKNOWN"
	config.UnknownIndexValue = 0

	return config
}
