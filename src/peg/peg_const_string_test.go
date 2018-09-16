package peg

import (
	"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestPegConstString(t *testing.T) {
	testdata := []struct {
		value           string
		name            string
		isCaseSensitive bool
		str             string
	}{
		{"", "", true, ""},
		{"a", "", false, "\"a\""},
		{"test", "", true, "'test'"},
		{"a-b", "", false, "\"a-b\""},
		{"a-b", "test", true, "test"},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			expr := NewPegConstString(v.name, []byte(v.value), v.isCaseSensitive)
			str := expr.String()

			test.EXPECT_EQ(t, str, v.str, "")
		})
	}
}
