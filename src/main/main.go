package main

import (
	//"charset"
	//"fmt"
	//"io"
	//"os"
	//"reflect"

	//"github.com/lioneagle/abnf/src/backends/c"
	//"github.com/lioneagle/abnf/src/charset"
	//"github.com/lioneagle/abnf/src/gen/charset_gen"
	"github.com/lioneagle/abnf/src/backends/golang"
	"github.com/lioneagle/abnf/src/gen/key_gen"
	"github.com/lioneagle/abnf/src/keys"
)

type A struct {
	x int
}

func (a *A) f(int) bool {
	return false
}

func main() {
	genSipMethod()
	genSipHeader()

	/*
		config := charset_gen.NewConfig()

		config.SetMaskPrefix("PS_SIP_CHARSETS")
		config.SetActionPrefix("PS_SIP")
		config.VarTypeName = "PS_DWORD"
		config.SetVarTypeSize(4)
		config.SetVarName("g_sipCharsets")
		config.ActionFirstLower = true
		config.UseBit = true

		charsets := charset_gen.NewCharsetTable()

		info := charset_gen.NewCharsetInfo("digit")
		info.Charset = charset.NewCharset()
		info.Charset.UniteRange(&charset.Range{'0', '9' + 1})
		charsets.Add(info)

		info = charset_gen.NewCharsetInfo("alpha")
		info.Charset = charset.NewCharset()
		info.Charset.UniteRange(&charset.Range{'a', 'z' + 1})
		info.Charset.UniteRange(&charset.Range{'A', 'Z' + 1})
		charsets.Add(info)

		charsets.Calc(config)

		gen_c := c.NewCharsetTableGeneratorForC()
		gen_c.GenerateFile(config, charsets, "ps_sip_charsets_1", ".")
	*/

	/*
		var gen charset.CharsetGenForCpp

		gen.GenerateMask(os.Stdout, nil)

		var r1 charset.Range

		//fmt.Printf("0x%x\n", uint32(-1))

		r1 = charset.Range{1, 2}

		r1.PrintAsChar(os.Stdout).WriteString("\n")

		r1 = charset.Range{1, 6}
		r1.PrintAsChar(os.Stdout).WriteString("\n")

		r1 = charset.Range{0, 257}
		r1.Print(os.Stdout).WriteString("\n")
		r1.PrintEachChar(os.Stdout).WriteString("\n")

		fmt.Println("r1 = ", r1)
		fmt.Printf("%c\n", '\\')

		var a A

		p := a.f

		fmt.Println("type =", reflect.TypeOf(p))
	*/
}

func genSipHeader() {
	config := key_gen.NewConfig()

	config.ActionName = "GetSipHeaderIndex"

	config.UnknownIndexName = "SIP_HDR_UNKNOWN"
	config.UnknownIndexValue = 0
	config.MaxIndexName = "SIP_HDR_MAX_NUM"

	name := "SipHeaderIndex"

	config.BraceAtNextLine = false
	config.CaseSensitive = false
	config.UseTabIndent = true
	config.PackageName = "sipparser"
	config.CursorName = "pos"
	config.CursorTypeName = "AbnfPos"
	config.IndexTypeName = "SipHeaderIndexType"
	config.IndexTypeSize = 4
	config.CharsetEnabled = true
	config.CharsetName = "IsSipToken"
	config.GenVersion = true

	keys := buildKeysForSipHeader(config)

	gen_go := golang.NewKeyCmpGeneratorForGolang()
	gen_go.GenerateFile(config, keys, name, ".")
}

func buildKeysForSipHeader(config *key_gen.Config) *keys.Keys {
	data := []struct {
		Name       string
		IndexName  string
		IndexValue int
	}{
		{"Accpet", "SIP_HDR_ACCEPT", 1},                                               // RFC3261
		{"Accpet-Encoding", "SIP_HDR_ACCEPT_ENCODING", 2},                             // RFC3261
		{"Accpet-Language", "SIP_HDR_ACCEPT_ALNGUAGE", 3},                             // RFC3261
		{"Alert-Info", "SIP_HDR_ALERT_INFO", 4},                                       // RFC3261
		{"Allow", "SIP_HDR_ALLOW", 5},                                                 // RFC3261
		{"Authentication-Info", "SIP_HDR_AUTHENTICATION_INFO", 6},                     // RFC3261
		{"Authorization", "SIP_HDR_AUTHORIZATION", 7},                                 // RFC3261
		{"Call-ID", "SIP_HDR_CALL_ID", 8},                                             // RFC3261
		{"i", "SIP_HDR_CALL_ID", 8},                                                   // RFC3261
		{"Call-Info", "SIP_HDR_CALL_INFO", 9},                                         // RFC3261
		{"Contact", "SIP_HDR_CONTACT", 10},                                            // RFC3261
		{"m", "SIP_HDR_CONTACT", 10},                                                  // RFC3261
		{"Content-Disposition", "SIP_HDR_CONTENT_DISPOSITION", 11},                    // RFC3261
		{"Content-Encoding", "SIP_HDR_CONTENT_ENCODING", 12},                          // RFC3261
		{"e", "SIP_HDR_CONTENT_ENCODING", 12},                                         // RFC3261
		{"Content-Language", "SIP_HDR_CONTENT_LANGUAGE", 13},                          // RFC3261
		{"Content-Length", "SIP_HDR_CONTENT_LENGTH", 14},                              // RFC3261
		{"l", "SIP_HDR_CONTENT_LENGTH", 14},                                           // RFC3261
		{"Content-Type", "SIP_HDR_CONTENT_TYPE", 15},                                  // RFC3261
		{"c", "SIP_HDR_CONTENT_TYPE", 15},                                             // RFC3261
		{"CSeq", "SIP_HDR_CSEQ", 16},                                                  // RFC3261
		{"Date", "SIP_HDR_DATE", 17},                                                  // RFC3261
		{"Error-Info", "SIP_HDR_ERROR_INFO", 18},                                      // RFC3261
		{"Expires", "SIP_HDR_EXPIRES", 19},                                            // RFC3261
		{"From", "SIP_HDR_FROM", 20},                                                  // RFC3261
		{"f", "SIP_HDR_FROM", 20},                                                     // RFC3261
		{"In-Reply-To", "SIP_HDR_IN_REPLY_TO", 21},                                    // RFC3261
		{"Max-Forwards", "SIP_HDR_MAX_FORWARDS", 22},                                  // RFC3261
		{"Min-Expires", "SIP_HDR_MIN_EXPIRES", 23},                                    // RFC3261
		{"MIME-Version", "SIP_HDR_MIME_VERSION", 24},                                  // RFC3261
		{"Organization", "SIP_HDR_ORGANIZATION", 25},                                  // RFC3261
		{"Priority", "SIP_HDR_PRIORITY", 26},                                          // RFC3261
		{"Proxy-Authenticate", "SIP_HDR_PROXY_AUTHENTICATE", 27},                      // RFC3261
		{"Proxy-Authorization", "SIP_HDR_PROXY_AUTHORIZATION", 28},                    // RFC3261
		{"Proxy-Require", "SIP_HDR_PROXY_REQUIRE", 29},                                // RFC3261
		{"Record-Route", "SIP_HDR_RECORD_ROUTE", 30},                                  // RFC3261
		{"Reply-To", "SIP_HDR_REPLY_TO", 31},                                          // RFC3261
		{"Require", "SIP_HDR_REQUIRE", 32},                                            // RFC3261
		{"Retry-After", "SIP_HDR_RETRY_AFTER", 33},                                    // RFC3261
		{"Route", "SIP_HDR_ROUTE", 34},                                                // RFC3261
		{"Server", "SIP_HDR_SERVER", 35},                                              // RFC3261
		{"Subject", "SIP_HDR_SUBJECT", 36},                                            // RFC3261
		{"s", "SIP_HDR_SUBJECT", 36},                                                  // RFC3261
		{"Supported", "SIP_HDR_SUPPORTED", 37},                                        // RFC3261
		{"k", "SIP_HDR_SUPPORTED", 37},                                                // RFC3261
		{"Timestamp", "SIP_HDR_TIMESTAMP", 38},                                        // RFC3261
		{"To", "SIP_HDR_TO", 39},                                                      // RFC3261
		{"t", "SIP_HDR_TO", 39},                                                       // RFC3261
		{"Unsupported", "SIP_HDR_UNSUPPORTED", 40},                                    // RFC3261
		{"User-Agent", "SIP_HDR_USER_AGENT", 41},                                      // RFC3261
		{"Via", "SIP_HDR_VIA", 42},                                                    // RFC3261
		{"v", "SIP_HDR_VIA", 42},                                                      // RFC3261
		{"Warning", "SIP_HDR_WARNING", 43},                                            // RFC3261
		{"WWW-Authenticate", "SIP_HDR_WWW_AUTHENTICATE", 44},                          // RFC3261
		{"RSeq", "SIP_HDR_RSEQ", 45},                                                  // RFC3262 (PRACK)
		{"RAck", "SIP_HDR_RACK", 46},                                                  // RFC3262 (PRACK)
		{"Subscription-State", "SIP_HDR_SUBSCRIPTION_STATE", 47},                      // RFC3265/RFC6665 (SUBSCRIBE and NOTIFY)
		{"Allow-Events", "SIP_HDR_ALLOW_EVENTS", 48},                                  // RFC3265/RFC6665 (SUBSCRIBE and NOTIFY)
		{"u", "SIP_HDR_ALLOW_EVENTS", 48},                                             // RFC3265/RFC6665 (SUBSCRIBE and NOTIFY)
		{"Event", "SIP_HDR_EVENT", 49},                                                // RFC3265/RFC6665 (SUBSCRIBE and NOTIFY)
		{"o", "SIP_HDR_EVENT", 49},                                                    // RFC3265/RFC6665 (SUBSCRIBE and NOTIFY)
		{"P-Media-Authorization", "SIP_HDR_P_MEDIA_AUTHORIZATION", 50},                // RFC3313 (Private Mechanism for SIP)
		{"Privacy", "SIP_HDR_PRIVACY", 51},                                            // RFC3323 (Privacy Mechanism)
		{"P-Asserted-Identity", "SIP_HDR_P_ASSERTED_IDENTITY", 52},                    // RFC3325 (Asserted Identity)
		{"P-Preferred-Identity", "SIP_HDR_P_PREFERRED_IDENTITY", 53},                  // RFC3325 (Asserted Identity)
		{"Reason", "SIP_HDR_REASON", 54},                                              // RFC3326 (Reason Header for SIP)
		{"Path", "SIP_HDR_PATH", 55},                                                  // RFC3327 (Extension Header for SIP Registering No-Adjacent Contacts)
		{"Security-Client", "SIP_HDR_SECURITY_CLIENT", 56},                            // RFC3329 (Security Mechanism Agreement for SIP)
		{"Security-Server", "SIP_HDR_SECURITY_SERVER", 57},                            // RFC3329 (Security Mechanism Agreement for SIP)
		{"Security-Verify", "SIP_HDR_SECURITY_VERIFY", 58},                            // RFC3329 (Security Mechanism Agreement for SIP)
		{"P-Associated-URI", "SIP_HDR_P_ASSOCIATED_URI", 59},                          // RFC3455/RFC7315 (Private Header (P-Header) Extensions to SIP for 3GPP)
		{"P-Called-Party-ID", "SIP_HDR_P_CALLED_PARTY_ID", 60},                        // RFC3455/RFC7315 (Private Header (P-Header) Extensions to SIP for 3GPP)
		{"P-Visited-Network-Info", "SIP_HDR_P_VISITED_NETWORK_ID", 61},                // RFC3455/RFC7315 (Private Header (P-Header) Extensions to SIP for 3GPP)
		{"P-Access-Network-ID", "SIP_HDR_P_ACCESS_NETWORK_INFO", 62},                  // RFC3455/RFC7315 (Private Header (P-Header) Extensions to SIP for 3GPP)
		{"P-Charging-Function-Address", "SIP_HDR_P_CHARGING_FUNCTION_ADDRESSES", 63},  // RFC3455/RFC7315 (Private Header (P-Header) Extensions to SIP for 3GPP)
		{"P-Charging-Vector", "SIP_HDR_P_CHARGING_VECTOR", 64},                        // RFC3455/RFC7315 (Private Header (P-Header) Extensions to SIP for 3GPP)
		{"Refer-To", "SIP_HDR_REFER_TO", 65},                                          // RFC3515/RFC4508 (REFER)
		{"r", "SIP_HDR_REFER_TO", 65},                                                 // RFC3515/RFC4508 (REFER)
		{"P-DCS-Trace-Party-ID", "SIP_HDR_P_DCS_TRACE_PARTY_ID", 66},                  // RFC3603/RFC5503 (Private SIP Proxy-to-Proxy Extensions for Supporting the PacketCable Distributed Call Signaling)
		{"P-DCS-OSPS", "SIP_HDR_P_DCS_OSPS", 67},                                      // RFC3603/RFC5503 (Private SIP Proxy-to-Proxy Extensions for Supporting the PacketCable Distributed Call Signaling)
		{"P-DCS-Billing-Info", "SIP_HDR_P_DCS_BILLING_INFO", 68},                      // RFC3603/RFC5503 (Private SIP Proxy-to-Proxy Extensions for Supporting the PacketCable Distributed Call Signaling)
		{"P-DCS-LAES", "SIP_HDR_P_DCS_LAES", 69},                                      // RFC3603/RFC5503 (Private SIP Proxy-to-Proxy Extensions for Supporting the PacketCable Distributed Call Signaling)
		{"P-DCS-Redirect", "SIP_HDR_DCS_REDIRECT", 70},                                // RFC3603/RFC5503 (Private SIP Proxy-to-Proxy Extensions for Supporting the PacketCable Distributed Call Signaling)
		{"Service-Route", "SIP_HDR_SERVICE_ROUTE", 71},                                // RFC3608 (SIP Extension Header Field for Service Route Discovery During Registration)
		{"Accept-Contact", "SIP_HDR_ACCEPT_CONTACT", 72},                              // RFC3841 (Caller Preferences)
		{"a", "SIP_HDR_ACCEPT_CONTACT", 72},                                           // RFC3841 (Caller Preferences)
		{"Reject-Contact", "SIP_HDR_REJECT_CONTACT", 73},                              // RFC3841 (Caller Preferences)
		{"j", "SIP_HDR_REJECT_CONTACT", 73},                                           // RFC3841 (Caller Preferences)
		{"Request-Disposition", "SIP_HDR_REQUEST_DISPOSITION", 74},                    // RFC3841 (Caller Preferences)
		{"Replaces", "SIP_HDR_REPLACES", 75},                                          // RFC3891 (SIP Replaces Header)
		{"Referred-By", "SIP_HDR_REFERRED_BY", 76},                                    // RFC3892 (SIP Referred-By Mechanism)
		{"b", "SIP_HDR_REFERRED_BY", 76},                                              // RFC3892 (SIP Referred-By Mechanism)
		{"SIP-ETag", "SIP_HDR_SIP_ETAG", 77},                                          // RFC3903 (SIP Extension for Event State)
		{"SIP-If-Match", "SIP_SIP_IF_MATCH", 78},                                      // RFC3903 (SIP Extension for Event State)
		{"Join", "SIP_HDR_JOIN", 79},                                                  // RFC3911 (SIP Join Header)
		{"Content-Transfer-Encoding", "SIP_HDR_CONTENT_TRANSFER_ENCODING", 80},        // RFC2045 and RFC2046 and RFC0822 (for MIME)
		{"Content-ID", "SIP_HDR_CONTENT_ID", 81},                                      // RFC2045/RFC8262 and RFC2046 and RFC0822 (for MIME)
		{"Content-Description", "SIP_HDR_CONTENT_DESCRIPTION", 82},                    // RFC2045 and RFC2046 and RFC0822 (for MIME)
		{"Session-Expires", "SIP_HDR_SESSION_EXPIRES", 83},                            // RFC4028 (Session Timer)
		{"x", "SIP_HDR_SESSION_EXPIRES", 83},                                          // RFC4028 (Session Timer)
		{"Min-SE", "SIP_HDR_MIN_SE", 84},                                              // RFC4028 (Session Timer)
		{"History-Info", "SIP_HDR_HISTORY_INFO", 85},                                  // RFC4244/RFC7044/RFC7544 (Request History Information)
		{"Resource-Priority", "SIP_HDR_RESOURCE_PRIORITY", 86},                        // RFC4412 (Communications Resource Priority for SIP)
		{"Accept-Resource-Priority", "SIP_HDR_ACCEPT_RESOURCE_PRIORITY", 87},          // RFC4412 (Communications Resource Priority for SIP)
		{"P-User-Database", "SIP_HDR_P_USER_DATABASE", 88},                            // RFC4457 (SIP P-User-Database header)
		{"Identity", "SIP_HDR_IDENTITY", 89},                                          // RFC4474/RFC8224 (Enhancements for Authenticated Identity Management in SIP)
		{"Identity-Info", "SIP_HDR_IDENTITY_INFO", 90},                                // RFC4474 (Enhancements for Authenticated Identity Management in SIP)
		{"Refer-Sub", "SIP_HDR_REFER_SUB", 91},                                        // RFC4488 (Suppression of SIP REFER Method Implicit Subscription)
		{"Target-Dialog", "SIP_HDR_TARGET_DIALOG", 92},                                // RFC4538 (Request Authorization through Dialog Id in SIP)
		{"P-Answer-State", "SIP_HDR_P_ANSWER_STATE", 93},                              // RFC4964  (The P-Answer-State Header Extension to SIP for the Open Mobile Alliance Push to Talk over Cellular)
		{"P-Profile-Key", "SIP_HDR_P_PROFILE_KEY", 94},                                // RFC5002  (P-Profile-Key Private Header)
		{"P-Early-Media", "SIP_HDR_P_EARLY_MEDIA", 95},                                // RFC5009 (P-Early-Media)
		{"P-Refused-URI-List", "SIP_HDR_P_REFUSED_URI_LIST", 96},                      // RFC5318  (SIP P-Refused-URI-List Private-Header)
		{"Permission-Missing", "SIP_HDR_PERMISSION_MISSING", 97},                      // RFC5360  (A Framework for Consent-Based Communications in SIP)
		{"Trigger-Consent", "SIP_HDR_TRIGGER_CONSENT", 98},                            // RFC5360  (A Framework for Consent-Based Communications in SIP)
		{"Answer-Mode", "SIP_HDR_ANSWER_MODE", 99},                                    // RFC5373  (Requesting Answering Modes for SIP)
		{"Priv-Answer-Mode", "SIP_HDR_PRIV_ANSWER_MODE", 100},                         // RFC5373  (Requesting Answering Modes for SIP)
		{"Max-Breadth", "SIP_HDR_MAX_BREADTH", 101},                                   // RFC5393  (Addressing an Amplification Vulnerability in SIP Forking Proxies)
		{"P-Served-User", "SIP_HDR_P_SERVED_USER", 102},                               // RFC5502  (SIP P-Served-User Private-Header for 3GPP IMS)
		{"Flow-Timer", "SIP_HDR_FLOW_TIMER", 103},                                     // RFC5626  (Managing Client-Initiated Connections in SIP)
		{"Suppress-If-Match", "SIP_HDR_SUPPRESS_IF_MATCH", 104},                       // RFC5839  (A Extension to SIP Events for Conditional Event Notification)
		{"Diversion", "SIP_HDR_DIVERSION", 105},                                       // RFC6044/RFC7544  (Mapping and Interworking of Diversion Information between Diversion and History-Info Headers in SIP)
		{"P-Asserted-Service", "SIP_HDR_P_ASSERTED_SERVICE", 106},                     // RFC6050  (SIP Extension for the Identification of Services)
		{"P-Preferred-Service", "SIP_HDR_P_PREFERRED_SERVICE", 107},                   // RFC6050  (SIP Extension for the Identification of Services)
		{"Info-Package", "SIP_HDR_INFO_PACKAGE", 108},                                 // RFC6086  (INFO Method and Package Framework)
		{"Recv-Info", "SIP_HDR_RECV_INFO", 109},                                       // RFC6086  (INFO Method and Package Framework)
		{"Geolocation", "SIP_HDR_GEOLOCATION", 110},                                   // RFC6442  (Location Conveyance for the Session Initiation Protocol)
		{"Geolocation-Routing", "SIP_HDR_GEOLOCATION_ROUTING", 111},                   // RFC6442  (Location Conveyance for the Session Initiation Protocol)
		{"Geolocation-Error", "SIP_HDR_GEOLOCATION_ERROR", 112},                       // RFC6442  (Location Conveyance for the Session Initiation Protocol)
		{"Policy-ID", "SIP_HDR_POLICY_ID", 113},                                       // RFC6794  (Framework for SIP Session Policies)
		{"Policy-Contact", "SIP_HDR_POLICY_CONTACT", 114},                             // RFC6794  (Framework for SIP Session Policies)
		{"Feature-Caps", "SIP_HDR_FEATURE_CAPS", 115},                                 // RFC6809  (Mechanism to Indicate Support of Features and Capabilities in SIP)
		{"P-Private-Network-Indication", "SIP_HDR_P_PRIVATE_NETWORK_INDICATION", 116}, // RFC7316  (P-Private-Network-Indication Private Header)
		{"Session-ID", "SIP_HDR_SESSION_ID", 117},                                     // RFC7329/RFC7989  (Session Identifier for SIP)
		{"User-to-User", "SIP_HDR_USER_TO_USER", 118},                                 // RFC7433  (A Mechanism for Transporting User-to-User Call Control Information in SIP)
		{"Refer-Events-At", "SIP_HDR_Refer_Events_At", 119},                           // RFC7614  (Explicit Subscriptions for the REFER Method)
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

func genSipMethod() {
	config := key_gen.NewConfig()

	config.ActionName = "GetSipMethodIndex"

	config.UnknownIndexName = "ABNF_SIP_METHOD_UNKNOWN"
	config.UnknownIndexValue = 0

	name := "SipMethodIndex"

	config.BraceAtNextLine = false
	config.CaseSensitive = true
	config.UseTabIndent = true
	config.PackageName = "sipparser"
	config.CursorName = "pos"
	config.CursorTypeName = "AbnfPos"
	config.IndexTypeName = "byte"
	config.IndexTypeSize = 1
	config.CharsetEnabled = true
	config.CharsetName = "IsSipToken"
	config.GenVersion = true

	keys := buildKeysForSipMethod(config)

	gen_go := golang.NewKeyCmpGeneratorForGolang()
	gen_go.GenerateFile(config, keys, name, ".")

}

func buildKeysForSipMethod(config *key_gen.Config) *keys.Keys {
	data := []struct {
		Name       string
		IndexName  string
		IndexValue int
	}{
		{"INVITE", "ABNF_SIP_METHOD_INVITE", 1},
		{"PRACK", "ABNF_SIP_METHOD_PRACK", 2},
		{"UPDATE", "ABNF_SIP_METHOD_UPDATE", 3},
		{"INFO", "ABNF_SIP_METHOD_INFO", 4},
		{"ACK", "ABNF_SIP_METHOD_ACK", 5},
		{"BYE", "ABNF_SIP_METHOD_BYE", 6},
		{"REGISTER", "ABNF_SIP_METHOD_REGISTER", 7},
		{"SUBSCRIBE", "ABNF_SIP_METHOD_SUBSCRIBE", 8},
		{"NOTIFY", "ABNF_SIP_METHOD_NOTIFY", 9},
		{"REFER", "ABNF_SIP_METHOD_REFER", 10},
		{"OPTIONS", "ABNF_SIP_METHOD_OPTIONS", 11},
		{"MESSAGE", "ABNF_SIP_METHOD_MESSAGE", 12},
		{"PUBLISH", "ABNF_SIP_METHOD_PUBLISH", 13},
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
