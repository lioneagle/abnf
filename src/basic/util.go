package basic

import (
	"fmt"
	"io"
	"strconv"
)

func PrintIndent(w io.Writer, indent int) {
	fmt.Fprintf(w, fmt.Sprintf("%%%ds", indent), "")
}

func PrintIntAsChar(w io.Writer, ch int32) io.Writer {
	switch ch {
	case '\a':
		fmt.Fprint(w, "\\a")
	case '\b':
		fmt.Fprint(w, "\\b")
	case '\f':
		fmt.Fprint(w, "\\f")
	case '\n':
		fmt.Fprint(w, "\\n")
	case '\r':
		fmt.Fprint(w, "\\r")
	case '\t':
		fmt.Fprint(w, "\\t")
	case '\v':
		fmt.Fprint(w, "\\v")
	case '\\':
		fmt.Fprint(w, "\\\\")
	case '"':
		fmt.Fprint(w, "\\\"")
	case '\'':
		fmt.Fprint(w, "\\'")
	case '-':
		fmt.Fprint(w, "\\-")
	case ' ':
		fmt.Fprint(w, "' '")
	default:
		if ch >= 0 && ch < 256 {
			if strconv.IsPrint(ch) && ch <= '~' {
				fmt.Fprintf(w, "'%c'", ch)
			} else {
				fmt.Fprintf(w, "\\x%02x", ch)
			}
		} else {
			fmt.Fprintf(w, "%d", ch)
		}
	}

	return w
}

func hexToInt(hex byte) int32 {
	if hex >= '0' && hex <= '9' {
		return int32(hex) - '0'
	}

	if hex >= 'a' && hex <= 'f' {
		return int32(hex) - 'a' + 10
	}

	if hex >= 'A' && hex <= 'F' {
		return int32(hex) - 'A' + 10
	}

	return -1
}

func octToInt(hex byte) int32 {
	if hex >= '0' && hex <= '7' {
		return int32(hex) - '0'
	}

	return -1
}

func UnescapeChar(src []byte, pos int) (ch int32, newPos int) {
	newPos = pos
	if newPos >= len(src) {
		return -1, newPos
	}

	if src[newPos] != '\\' {
		return int32(src[newPos]), newPos + 1
	}

	newPos++

	if newPos >= len(src) {
		return '\\', newPos
	}

	ch1 := src[newPos]

	if ch1 == 'x' {
		return unescapeHex(src, newPos)
	}

	if ch1 >= '0' && ch1 <= '3' {
		return unescapeOct(src, newPos)
	}

	//fmt.Printf("ch1 = %d\n", ch1)
	//fmt.Printf("dst = %d\n", 'a')

	switch ch1 {
	case 'a':
		return '\a', newPos + 1
	case 'b':
		return '\b', newPos + 1
	case 'f':
		return '\f', newPos + 1
	case 'n':
		return '\n', newPos + 1
	case 'r':
		return '\r', newPos + 1
	case 't':
		return '\t', newPos + 1
	case 'v':
		return '\v', newPos + 1
	case '\\':
		return '\\', newPos + 1
	case '"':
		return '"', newPos + 1
	case '\'':
		return '\'', newPos + 1
	case '-':
		return '-', newPos + 1
	default:
		return '\\', newPos
	}

}

func unescapeHex(src []byte, pos int) (ch int32, newPos int) {
	newPos = pos
	if (newPos + 2) >= len(src) {
		return '\\', newPos
	}

	high := hexToInt(src[newPos+1])
	low := hexToInt(src[newPos+2])

	if low < 0 || high < 0 {
		return '\\', newPos
	}

	return (high << 4) | low, newPos + 3
}

func unescapeOct(src []byte, pos int) (ch int32, newPos int) {
	newPos = pos
	if (newPos + 2) >= len(src) {
		return '\\', newPos
	}

	c1 := int32(src[newPos]) - '0'
	c2 := octToInt(src[newPos+1])
	c3 := octToInt(src[newPos+2])

	if c2 < 0 || c3 < 0 {
		return '\\', newPos
	}

	return (c1 << 6) | (c2 << 3) | c3, newPos + 3
}

func RoundToAlign(x, align int) int {
	return (x + align - 1) & ^(align - 1)
}
