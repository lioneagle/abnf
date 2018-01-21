package basic

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestPrintIdent(t *testing.T) {
	var buf bytes.Buffer

	PrintIndent(&buf, 2)
	str := buf.String()
	test.EXPECT_EQ(t, str, "  ", "")
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

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			var buf bytes.Buffer
			PrintIntAsChar(&buf, v.src)
			str := buf.String()
			test.EXPECT_EQ(t, str, v.wanted, "")
		})
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

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			ch, newPos := UnescapeChar([]byte(v.src), 0)
			test.EXPECT_EQ(t, ch, v.wanted, "")
			test.EXPECT_EQ(t, newPos, v.newPos, "")
		})
	}
}
