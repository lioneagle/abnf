package charset

import (
	"bytes"
	//"os"
	//"strconv"
	"testing"
)

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
