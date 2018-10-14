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
		{"To", "ABNF_SIP_HDR_TO", 2},
		{"t", "ABNF_SIP_HDR_TO", 2},
		{"Via", "ABNF_SIP_HDR_VIA", 3},
		{"v", "ABNF_SIP_HDR_VIA", 3},
		{"Call-ID", "ABNF_SIP_HDR_CALL_ID", 4},
		{"i", "ABNF_SIP_HDR_CALL_ID", 4},
		{"CSeq", "ABNF_SIP_HDR_CSEQ", 5},
		{"Content-Length", "ABNF_SIP_HDR_CONTENT_LENGTH", 6},
		{"l", "ABNF_SIP_HDR_CONTENT_LENGTH", 6},
		{"Content-Type", "ABNF_SIP_HDR_CONTENT_TYPE", 7},
		{"c", "ABNF_SIP_HDR_CONTENT_TYPE", 7},
		{"Contact", "ABNF_SIP_HDR_CONTACT", 8},
		{"m", "ABNF_SIP_HDR_CONTACT", 8},
		{"Max-Forwards", "ABNF_SIP_HDR_MAX_FORWARDS", 9},
		{"Route", "ABNF_SIP_HDR_ROUTE", 10},
		{"Record-Route", "ABNF_SIP_HDR_RECORD_ROUTE", 11},
		{"Content-Disposition", "ABNF_SIP_HDR_CONTENT_DISPOSITION", 12},
		{"Allow", "ABNF_SIP_HDR_ALLOW", 13},
		{"Content-Encoding", "ABNF_SIP_HDR_CONTENT_ENCODING", 14},
		{"e", "ABNF_SIP_HDR_CONTENT_ENCODING", 14},
		{"Date", "ABNF_SIP_HDR_DATE", 15},
		{"Subject", "ABNF_SIP_HDR_SUBJECT", 16},
		{"s", "ABNF_SIP_HDR_SUBJECT", 16},
		{"Supported", "ABNF_SIP_HDR_SUPPORTED", 17},
		{"k", "ABNF_SIP_HDR_SUPPORTED", 17},
		{"Allow-Events", "ABNF_SIP_HDR_ALLOW_EVENTS", 18},
		{"u", "ABNF_SIP_HDR_ALLOW_EVENTS", 18},
		{"Event", "ABNF_SIP_HDR_EVENT", 19},
		{"o", "ABNF_SIP_HDR_EVENT", 19},
		{"Refer-To", "ABNF_SIP_HDR_REFER_TO", 20},
		{"r", "ABNF_SIP_HDR_REFER_TO", 20},
		{"Accept-Contact", "ABNF_SIP_HDR_ACCEPT_CONTACT", 21},
		{"a", "ABNF_SIP_HDR_ACCEPT_CONTACT", 21},
		{"Reject-Contact", "ABNF_SIP_HDR_REJECT_CONTACT", 22},
		{"j", "ABNF_SIP_HDR_REJECT_CONTACT", 22},
		{"Request-Disposition", "ABNF_SIP_HDR_REQUEST_DISPOSITION", 23},
		{"Referred-By", "ABNF_SIP_HDR_REFERRED_BY", 24},
		{"b", "ABNF_SIP_HDR_REFERRED_BY", 24},
		{"Session-Expires", "ABNF_SIP_HDR_SESSION_EXPIRES", 25},
		{"x", "ABNF_SIP_HDR_SESSION_EXPIRES", 25},
		{"MIME-Version", "ABNF_SIP_HDR_MIME_VERSION", 26},
	}

	ret := keys.NewKeys()

	for _, v := range data {
		key := &keys.Key{Name: v.Name, Index: keys.Index{Name: v.IndexName, Value: v.IndexValue}}
		ret.Add(key)
	}

	if len(config.UnknownIndexName) > 0 {
		ret.AddIndex(&keys.Index{Name: config.UnknownIndexName, Value: config.UnknownIndexValue})
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
