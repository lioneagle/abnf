package basic

import (
	"bytes"
	//"fmt"
	"testing"
	"trace"
)

func TestPrintIdent(t *testing.T) {
	prefix := trace.CallerName(0)
	var buf bytes.Buffer

	PrintIndent(&buf, 2)
	str := buf.String()
	if str != "  " {
		t.Errorf("%s failed: space-num = %d, wanted = 2\r\n", prefix, len(str))
	}
}

func TestPrintChar(t *testing.T) {
	testdata := []struct {
		src    int32
		wanted string
	}{
		{'\a', "\\a"},
		{'\b', "\\b"},
		{'\f', "\\f"},
		{'\n', "\\n"},
		{'\r', "\\r"},
		{'\t', "\\t"},
		{'\v', "\\v"},
		{'\\', "\\\\"},
		{'"', "\\\""},
		{'\'', "\\'"},
		{'-', "\\-"},

		{'a', "a"},
		{'~', "~"},
		{1, "\\x01"},
		{288, "288"},
		{-1, "-1"},
	}
	prefix := trace.CallerName(0)
	var buf bytes.Buffer

	for i, v := range testdata {
		buf.Reset()
		PrintIntAsChar(&buf, v.src)
		str := buf.String()
		if str != v.wanted {
			t.Errorf("%s[%d] failed: str = %s, wanted = %s\n", prefix, i, str, v.wanted)
		}
	}

}

func TestUnescapeChar(t *testing.T) {
	testdata := []struct {
		src    string
		wanted int32
		newPos int
	}{
		{"\\a", '\a', 2},
		{"\\b", '\b', 2},
		{"\\f", '\f', 2},
		{"\\n", '\n', 2},
		{"\\r", '\r', 2},
		{"\\t", '\t', 2},
		{"\\v", '\v', 2},
		{"\\\\", '\\', 2},
		{"\\\"", '"', 2},
		{"\\'", '\'', 2},
		{"\\-", '-', 2},

		{"\\", '\\', 1},

		{"a", 'a', 1},
		{"-", '-', 1},
		{"0", '0', 1},
		{"9", '9', 1},
		{"~", '~', 1},

		{"\\017", '\017', 4},
		{"\\123", '\123', 4},
		{"\\203", '\203', 4},
		{"\\377", '\377', 4},
		{"\\187", '\\', 1},
		{"\\477", '\\', 1},
		{"\\p77", '\\', 1},
		{"\\07", '\\', 1},
		{"\\1", '\\', 1},

		{"\\x12", 0x12, 4},
		{"\\x0F", 0x0f, 4},
		{"\\x9a", 0x9a, 4},
		{"\\xf0", 0xf0, 4},
		{"\\xA9", 0xa9, 4},
		{"\\xG9", '\\', 1},
		{"\\x9", '\\', 1},
		{"\\x", '\\', 1}, //*/
	}
	prefix := trace.CallerName(0)

	for i, v := range testdata {
		ch, newPos := UnescapeChar([]byte(v.src), 0)
		if ch != v.wanted {
			t.Errorf("%s[%d] failed: ch = %d, wanted = %d\n", prefix, i, ch, v.wanted)
		}

		if newPos != v.newPos {
			t.Errorf("%s[%d] failed: newPos = %d, wanted = %d\n", prefix, i, newPos, v.newPos)
		}
	}
}
