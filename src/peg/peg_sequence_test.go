package peg

import (
	"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestPegSequence(t *testing.T) {
	testdata := []struct {
		name  string
		expr1 Peg
		expr2 Peg
		str   string
	}{
		{"", NewPegCharset("", []byte("a")), NewPegCharset("", []byte("A-G")), "'a' ['A'-'G']"},
		{"", NewPegCharset("xyz", []byte("a")), NewPegCharset("", []byte("A-G")), "xyz ['A'-'G']"},
		{"abc", NewPegCharset("xyz", []byte("a")), NewPegCharset("", []byte("A-G")), "abc"},
		{"", NewPegCharset("alpha", []byte("a-zA-Z")), NewPegCharset("digit", []byte("0-9")), "alpha digit"},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			expr := NewPegSequence(v.name)
			expr.Append(v.expr1)
			expr.Append(v.expr2)
			str := expr.String()

			test.EXPECT_EQ(t, str, v.str, "")
		})
	}
}
