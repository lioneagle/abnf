package charset

import (
	//"bytes"
	//"os"
	//"strconv"
	"testing"
	"trace"
)

func TestRangeSize(t *testing.T) {
	testdata := []struct {
		r    Range
		size int32
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

/*func TestRangeContains(t *testing.T) {
	testdata := []struct {
		src    int32
		wanted string
	}{
		{int32('\a'), "\\a"},
		{int32('\b'), "\\b"},
		{int32('\f'), "\\f"},
		{int32('\n'), "\\n"},
		{int32('\r'), "\\r"},
		{int32('\t'), "\\t"},
		{int32('\v'), "\\v"},
		{int32('\\'), "\\\\"},
		{int32('"'), "\\\""},
		{int32('\''), "\\'"},
		{int32('-'), "\\-"},

		{int32('a'), "a"},
		{int32('~'), "~"},
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

}*/

/*
func checkPrintResult(t *testing.T, name string, f func(ByteAndStringWriter) ByteAndStringWriter, wanted string) {
	buf := bytes.NewBuffer(nil)
	f(buf)
	if buf.String() != wanted {
		t.Errorf("%s Print wrong fomat, ret = \"%s\", wanted = \"%s\"", name, buf.String(), wanted)
	}
}

func checkRangeContains(t *testing.T, r *Range, lowMin, highMax int32) {

	var i int32

	low := r.Low
	high := r.High

	for i := low; i < high; i++ {
		if !r.Contains(i) {
			t.Errorf("Range %v should contain %d", r, i)
		}
	}

	for i = lowMin; i < low; i++ {
		if r.Contains(i) {
			t.Errorf("Range %v should not contain %d", r, i)
		}
	}

	for i := high; i < highMax; i++ {
		if r.Contains(i) {
			t.Errorf("Range %v should not contain %d", r, i)
		}
	}
}

func TestRangeContains(t *testing.T) {
	checkRangeContains(t, &Range{10, 20}, 0, 256)
	checkRangeContains(t, &Range{10, 20}, -100, 300)
}

func TestRangePrint(t *testing.T) {
	checkPrintResult(t, "Range", (&Range{-1, 0}).PrintAsChar, "-1")
	checkPrintResult(t, "Range", (&Range{258, 259}).PrintAsChar, "258")
	checkPrintResult(t, "Range", (&Range{1, 2}).Print, "1")
	checkPrintResult(t, "Range", (&Range{0x27, 0x28}).PrintAsChar, "'")
	checkPrintResult(t, "Range", (&Range{0x5c, 0x5d}).PrintAsChar, "\\\\")
	checkPrintResult(t, "Range", (&Range{1, 5}).Print, "1-5")
	checkPrintResult(t, "Range", (&Range{1, 5}).PrintAsChar, "\\x01-\\x05")
	checkPrintResult(t, "Range", (&Range{1, 5}).PrintEachChar, "\\x01, \\x02, \\x03, \\x04")
	checkPrintResult(t, "Range", (&Range{7, 14}).PrintEachChar, "\\a, \\b, \\t, \\n, \\v, \\f, \\r")
	checkPrintResult(t, "Range", (&Range{'a', 'd'}).PrintEachChar, "a, b, c")
}

func TestRangeEqual(t *testing.T) {
	if !(&Range{-1, 0}).Equal(&Range{-1, 0}) {
		t.Errorf("Range Equal test failed")
	}

	if (&Range{-1, 0}).Equal(&Range{-1, 2}) {
		t.Errorf("Range NotEqual test failed")
	}

}

func TestRangeLess(t *testing.T) {
	if !(&Range{-1, 0}).Less(&Range{1, 2}) {
		t.Errorf("Range Less test failed")
	}

	if (&Range{-1, 0}).Less(&Range{-1, 2}) {
		t.Errorf("Range NotLess test failed")
	}

	if (&Range{-1, 0}).Less(&Range{-2, 2}) {
		t.Errorf("Range NotLess test failed")
	}

}

func TestRangeLessEqual(t *testing.T) {
	if !(&Range{-1, 0}).LessEqual(&Range{1, 2}) {
		t.Errorf("Range Less test failed")
	}

	if !(&Range{-1, 0}).LessEqual(&Range{-1, 2}) {
		t.Errorf("Range Less test failed")
	}

	if (&Range{-1, 0}).LessEqual(&Range{-2, 2}) {
		t.Errorf("Range NotLess test failed")
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
*/
