package charset

import (
	//"bytes"
	//"fmt"
	//"os"
	//"strconv"
	"testing"
	"trace"
)

func TestCharsetUniteRange1(t *testing.T) {
	testdata := []struct {
		r1   Range
		r2   Range
		str  string
		size uint32
	}{
		{Range{0, 0}, Range{}, "", 0},
		{Range{0, 0}, Range{1, 2}, "1", 1},
		{Range{1, 2}, Range{7, 9}, "1, 7-9", 3},
		{Range{7, 9}, Range{1, 2}, "1, 7-9", 3},
		{Range{7, 9}, Range{11, 15}, "7-9, 11-15", 6},
		{Range{7, 10}, Range{2, 12}, "2-12", 10},
		{Range{7, 10}, Range{2, 7}, "2-10", 8},
		{Range{7, 10}, Range{10, 13}, "7-13", 6},
		{Range{7, 10}, Range{8, 12}, "7-12", 5},
		{Range{7, 10}, Range{1, 9}, "1-10", 9},
	}
	prefix := trace.CallerName(0)
	c := Charset{}

	for i, v := range testdata {
		c.RemoveAll()
		c.UniteRange(&v.r1)
		c.UniteRange(&v.r2)
		str := c.StringAsInt()

		if str != v.str {
			t.Errorf("%s[%d] failed: str = %s, wanted = %s\n", prefix, i, str, v.str)
		}

		if c.Size() != v.size {
			t.Errorf("%s[%d] failed: size = %d, wanted = %d\n", prefix, i, c.Size(), v.size)
		}
	}
}

func TestCharsetUniteRange2(t *testing.T) {
	testdata := []struct {
		ranges []Range
		str    string
		size   uint32
	}{
		{[]Range{{1, 2}}, "1", 1},                                                                                        // empty
		{[]Range{{1, 2}, {1, 2}}, "1", 1},                                                                                // same node
		{[]Range{{1, 2}, {2, 4}}, "1-4", 3},                                                                              // concat
		{[]Range{{1, 2}, {-2, -1}, {-5, -3}, {3, 5}}, "-5--3, -2, 1, 3-5", 6},                                            // insert before
		{[]Range{{1, 2}, {7, 9}}, "1, 7-9", 3},                                                                           // append at tail without cross
		{[]Range{{1, 2}, {7, 9}, {3, 4}}, "1, 3, 7-9", 4},                                                                // insert at middle without cross
		{[]Range{{1, 2}, {7, 9}, {3, 4}, {6, 10}}, "1, 3, 6-10", 6},                                                      // insert at middle without cross
		{[]Range{{1, 2}, {7, 9}, {3, 4}, {6, 10}, {6, 9}}, "1, 3, 6-10", 6},                                              // contained in one node
		{[]Range{{1, 2}, {7, 9}, {3, 4}, {6, 10}, {6, 9}, {5, 8}}, "1, 3, 5-10", 7},                                      // low is in one node and high is greater than last node's high
		{[]Range{{1, 2}, {7, 9}, {3, 4}, {6, 10}, {6, 9}, {5, 8}, {15, 18}}, "1, 3, 5-10, 15-18", 10},                    // append at tail without cross
		{[]Range{{1, 2}, {7, 9}, {3, 4}, {6, 10}, {6, 9}, {5, 8}, {15, 18}, {20, 22}}, "1, 3, 5-10, 15-18, 20-22", 12},   // append at tail without cross
		{[]Range{{1, 2}, {7, 9}, {3, 4}, {6, 10}, {6, 9}, {5, 8}, {15, 18}, {20, 22}, {7, 19}}, "1, 3, 5-19, 20-22", 18}, // low is in one node and high is greater than other two node
		{[]Range{{1, 2}, {3, 22}}, "1, 3-22", 20},                                                                        // low is equal to one node's high
		{[]Range{{1, 2}, {3, 22}, {23, 25}}, "1, 3-22, 23-25", 22},                                                       // low is equal to one node's high
		{[]Range{{1, 2}, {3, 22}, {23, 25}, {22, 23}}, "1, 3-25", 23},                                                    // high is equal to one node's low
		{[]Range{{1, 2}, {3, 22}, {23, 25}, {22, 23}, {100, 102}}, "1, 3-25, 100-102", 25},                               // append at tail without cross
		{[]Range{{1, 2}, {3, 22}, {23, 25}, {22, 23}, {100, 102}, {40, 42}}, "1, 3-25, 40-42, 100-102", 27},              // low is less than one node's low and high is greater than onther node's high
	}
	prefix := trace.CallerName(0)

	for i, v := range testdata {
		c := &Charset{}
		c.UniteRangeSlice(v.ranges)
		str := c.StringAsInt()

		if str != v.str {
			t.Errorf("%s[%d] failed: str = %s, wanted = %s\n", prefix, i, str, v.str)
		}

		if c.Size() != v.size {
			t.Errorf("%s[%d] failed: size = %d, wanted = %d\n", prefix, i, c.Size(), v.size)
		}
	}
}

func TestCharsetContains(t *testing.T) {
	ranges := []Range{Range{1, 2}, {3, 22}, {23, 25}, {22, 23}, {100, 102}, {40, 42}}
	testdata := []struct {
		val      int32
		contains bool
	}{
		{1, true},
		{10, true},
		{24, true},

		{-1, false},
		{2, false},
		{102, false},
		{50, false},
	}
	prefix := trace.CallerName(0)

	c := &Charset{}
	c.UniteRangeSlice(ranges)

	for i, v := range testdata {
		if !c.Contains(v.val) && v.contains {
			t.Errorf("%s[%d] failed: should be contained\n", prefix, i)
		}

		if c.Contains(v.val) && !v.contains {
			t.Errorf("%s[%d] failed: should not be contained\n", prefix, i)
		}
	}
}

func TestCharsetUniteChar(t *testing.T) {
	testdata := []struct {
		ranges []Range
		ch     int32
		str    string
		size   uint32
	}{
		{[]Range{}, 1, "1", 1},
		{[]Range{{1, 3}}, 0, "0-3", 3},
		{[]Range{{1, 3}}, 3, "1-4", 3},
		{[]Range{{1, 3}, {4, 10}}, 3, "1-10", 9},
		{[]Range{{1, 3}, {4, 10}}, 7, "1-3, 4-10", 8},
		{[]Range{{1, 3}, {7, 10}}, 5, "1-3, 5, 7-10", 6},
	}
	prefix := trace.CallerName(0)

	for i, v := range testdata {
		c := &Charset{}
		c.UniteRangeSlice(v.ranges)
		c.UniteChar(v.ch)
		str := c.StringAsInt()

		if str != v.str {
			t.Errorf("%s[%d] failed: str = %s, wanted = %s\n", prefix, i, str, v.str)
		}

		if c.Size() != v.size {
			t.Errorf("%s[%d] failed: size = %d, wanted = %d\n", prefix, i, c.Size(), v.size)
		}
	}

}

func TestCharsetMakeFromBytes(t *testing.T) {
	testdata := []struct {
		src       string
		printType int
		str       string
		size      uint32
	}{
		{"", print_as_int, "", 0},
		{"\\x01\\002", print_as_int, "1-3", 2},
		{"\\x01\\002-\\x05", print_as_int, "1-5", 4},
		{"a-", print_as_char, "\\-, a", 2},
		{"\\21", print_as_char, "1-3, \\\\", 3},
		{"\\x05-\\x01", print_as_char, "\\x01-\\x05", 4},
		{"a-d", print_each_char, "a, b, c", 3},
	}
	prefix := trace.CallerName(0)

	for i, v := range testdata {
		c := &Charset{}
		c.MakeFromBytes([]byte(v.src))
		str := c.toString(v.printType)

		if str != v.str {
			t.Errorf("%s[%d] failed: str = %s, wanted = %s\n", prefix, i, str, v.str)
		}

		if c.Size() != v.size {
			t.Errorf("%s[%d] failed: size = %d, wanted = %d\n", prefix, i, c.Size(), v.size)
		}
	}
}
