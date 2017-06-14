package charset

import (
	"bytes"
	//"os"
	//"strconv"
	"testing"
	"trace"
)

func TestRangeSize(t *testing.T) {
	testdata := []struct {
		r    Range
		size uint32
	}{
		{Range{0, 0}, 0},
		{Range{1, 2}, 1},
		{Range{13, 40}, 27},
		{Range{-10, 2}, 12},
	}
	prefix := trace.CallerName(0)

	for i, v := range testdata {

		if v.r.Size() != v.size {
			t.Errorf("%s[%d] failed: size = %s, wanted = %s\n", prefix, i, v.r.Size(), v.size)
		}
	}

}

func TestRangeEqual(t *testing.T) {
	testdata := []struct {
		r1    Range
		r2    Range
		equal bool
	}{
		{Range{0, 0}, Range{}, true},
		{Range{1, 2}, Range{1, 2}, true},
		{Range{13, 40}, Range{13, 40}, true},
		{Range{-10, 2}, Range{-10, 2}, true},

		{Range{0, 0}, Range{1, 3}, false},
		{Range{1, 2}, Range{2, 2}, false},
		{Range{13, 40}, Range{-1, 40}, false},
		{Range{-10, 2}, Range{10, 2}, false},
	}
	prefix := trace.CallerName(0)

	for i, v := range testdata {
		if !v.r1.Equal(&v.r2) && v.equal {
			t.Errorf("%s[%d] failed: should be equal\n", prefix, i)
		}

		if v.r1.Equal(&v.r2) && !v.equal {
			t.Errorf("%s[%d] failed: should not be equal\n", prefix, i)
		}
	}
}

func TestRangeLess(t *testing.T) {
	testdata := []struct {
		r1   Range
		r2   Range
		less bool
	}{
		{Range{0, 0}, Range{1, 1}, true},
		{Range{1, 2}, Range{2, 3}, true},
		{Range{13, 40}, Range{14, 32}, true},
		{Range{-10, 2}, Range{-1, 3}, true},

		{Range{0, 0}, Range{-1, 3}, false},
		{Range{1, 2}, Range{-2, 2}, false},
		{Range{13, 40}, Range{7, 78}, false},
		{Range{-10, 2}, Range{-11, -3}, false},
	}
	prefix := trace.CallerName(0)

	for i, v := range testdata {
		if !v.r1.Less(&v.r2) && v.less {
			t.Errorf("%s[%d] failed: should be less\n", prefix, i)
		}

		if v.r1.Less(&v.r2) && !v.less {
			t.Errorf("%s[%d] failed: should not be less\n", prefix, i)
		}
	}
}

func TestRangeLessEqual(t *testing.T) {
	testdata := []struct {
		r1        Range
		r2        Range
		lessEqual bool
	}{
		{Range{0, 0}, Range{1, 1}, true},
		{Range{1, 2}, Range{2, 3}, true},
		{Range{13, 40}, Range{14, 32}, true},
		{Range{-10, 2}, Range{-1, 3}, true},
		{Range{0, 0}, Range{0, 1}, true},
		{Range{1, 2}, Range{1, 3}, true},
		{Range{13, 40}, Range{13, 32}, true},
		{Range{-10, 2}, Range{-10, 3}, true},

		{Range{0, 0}, Range{-1, 3}, false},
		{Range{1, 2}, Range{-2, 2}, false},
		{Range{13, 40}, Range{7, 78}, false},
		{Range{-10, 2}, Range{-11, -3}, false},
	}
	prefix := trace.CallerName(0)

	for i, v := range testdata {
		if !v.r1.LessEqual(&v.r2) && v.lessEqual {
			t.Errorf("%s[%d] failed: should be less-equal\n", prefix, i)
		}

		if v.r1.LessEqual(&v.r2) && !v.lessEqual {
			t.Errorf("%s[%d] failed: should not be less-equal\n", prefix, i)
		}
	}
}

func TestRangeAssert(t *testing.T) {

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("TestRangeAssert: should have panic")
		}
	}()

	(&Range{1, 0}).Assert()
}

func TestRangeContains(t *testing.T) {
	testdata := []struct {
		r        Range
		ch       int32
		contains bool
	}{
		{Range{0, 10}, 0, true},
		{Range{2, 10}, 9, true},
		{Range{-7, 10}, -7, true},

		{Range{1, 10}, 0, false},
		{Range{1, 10}, 10, false},
	}
	prefix := trace.CallerName(0)

	for i, v := range testdata {
		if !v.r.Contains(v.ch) && v.contains {
			t.Errorf("%s[%d] failed: should be contained\n", prefix, i)
		}

		if v.r.Contains(v.ch) && !v.contains {
			t.Errorf("%s[%d] failed: should not be contained\n", prefix, i)
		}
	}
}

func TestRangePrintAsInt(t *testing.T) {
	testdata := []struct {
		r   Range
		str string
	}{
		{Range{0, 11}, "0-10"},
		{Range{9, 10}, "9"},
		{Range{1, 1}, ""},
	}
	prefix := trace.CallerName(0)

	for i, v := range testdata {
		str := v.r.StringAsInt()
		if str != v.str {
			t.Errorf("%s[%d] failed: str = %s, wanted = %s\n", prefix, i, str, v.str)
		}
	}
}

func TestRangePrintAsChar(t *testing.T) {
	testdata := []struct {
		r   Range
		str string
	}{
		{Range{}, ""},
		{Range{'a', 'z' + 1}, "a-z"},
		{Range{'0', '9' + 1}, "0-9"},
		{Range{'\\', '\\' + 1}, "\\\\"},
		{Range{'\n', '\r' + 1}, "\\n-\\r"},
		{Range{1, 6}, "\\x01-\\x05"},
	}
	prefix := trace.CallerName(0)

	for i, v := range testdata {
		buf := &bytes.Buffer{}
		v.r.PrintAsChar(buf)
		str := buf.String()
		if str != v.str {
			t.Errorf("%s[%d] failed: str = %s, wanted = %s\n", prefix, i, str, v.str)
		}
	}
}

func TestRangePrintEachChar(t *testing.T) {
	testdata := []struct {
		r   Range
		str string
	}{
		{Range{}, ""},
		{Range{'a', 'd'}, "a, b, c"},
		{Range{'0', '3'}, "0, 1, 2"},
		{Range{'\\', '\\' + 1}, "\\\\"},
		{Range{'\n', '\r' + 1}, "\\n, \\v, \\f, \\r"},
		{Range{1, 5}, "\\x01, \\x02, \\x03, \\x04"},
	}
	prefix := trace.CallerName(0)

	for i, v := range testdata {
		buf := &bytes.Buffer{}
		v.r.PrintEachChar(buf)
		str := buf.String()
		if str != v.str {
			t.Errorf("%s[%d] failed: str = %s, wanted = %s\n", prefix, i, str, v.str)
		}
	}
}
