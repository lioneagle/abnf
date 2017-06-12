package charset

import (
	//"bytes"
	//"fmt"
	//"os"
	//"strconv"
	"testing"
	"trace"
)

func TestCharsetUniteRange2(t *testing.T) {
	testdata := []struct {
		ranges []Range
		str    string
		size   uint32
	}{
		{[]Range{{0, 0}, {}}, "", 0},
		{[]Range{{0, 0}, {1, 2}}, "1", 1},
		{[]Range{{1, 2}, {7, 9}}, "1, 7-9", 3},
		{[]Range{{7, 9}, {1, 2}}, "1, 7-9", 3},
		{[]Range{{7, 9}, {11, 15}}, "7-9, 11-15", 6},
		{[]Range{{7, 10}, {2, 12}}, "2-12", 10},
		{[]Range{{7, 10}, {2, 7}}, "2-10", 8},
		{[]Range{{7, 10}, {10, 13}}, "7-13", 6},
		{[]Range{{7, 10}, {8, 12}}, "7-12", 5},
		{[]Range{{7, 10}, {1, 9}}, "1-10", 9},

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

func TestCharsetRemoveAll(t *testing.T) {
	prefix := trace.CallerName(0)

	c := &Charset{}
	c.MakeFromBytes([]byte("\\x01-\\x05a-c\\k-m"))
	c.RemoveAll()
	str := c.StringAsInt()

	if str != "" {
		t.Errorf("%s failed: should be empty string\n", prefix)
	}

	if c.Size() != 0 {
		t.Errorf("%s failed: should be empty\n", prefix)
	}
}

func TestCharsetDifferenceRange(t *testing.T) {
	testdata := []struct {
		c    string
		r    Range
		str  string
		size uint32
	}{
		{"", Range{1, 2}, "", 0},
		{"\\x01-\\x02", Range{1, 2}, "", 0},
		{"\\x03-\\x0a", Range{2, 5}, "5-10", 5},
		{"\\x03-\\x0a", Range{2, 10}, "", 0},
		{"\\x03-\\x0a", Range{3, 11}, "", 0},
		{"\\x03-\\x0a", Range{4, 10}, "3", 1},
		{"\\x03-\\x0a", Range{5, 7}, "3-5, 7-10", 5},
		{"\\x03-\\x0a", Range{1, 3}, "3-10", 7},
		{"\\x03-\\x0a", Range{10, 12}, "3-10", 7},

		{"\\x01-\\x02\\x0a-\\x0f\\x1a-\\x1f", Range{1, 2}, "10-15, 26-31", 10},
		{"\\x01-\\x02\\x03-\\x0a\\x1a-\\x1f", Range{2, 5}, "1, 5-10, 26-31", 11},
		{"\\x01-\\x02\\x03-\\x0a\\x1a-\\x1f", Range{2, 10}, "1, 26-31", 6},
		{"\\x01-\\x02\\x03-\\x0a\\x1a-\\x1f", Range{3, 11}, "1, 26-31", 6},
		{"\\x01-\\x02\\x03-\\x0a\\x1a-\\x1f", Range{4, 10}, "1, 3, 26-31", 7},
		{"\\x01-\\x02\\x03-\\x0a\\x1a-\\x1f", Range{5, 7}, "1, 3-5, 7-10, 26-31", 11},
		{"\\x01-\\x02\\x03-\\x0a\\x1a-\\x1f", Range{1, 3}, "3-10, 26-31", 12},
		{"\\x01-\\x02\\x03-\\x0a\\x1a-\\x1f", Range{10, 12}, "1, 3-10, 26-31", 13},

		{"\\x01-\\x02\\x03-\\x0a\\x1a-\\x1f", Range{2, 28}, "1, 28-31", 4},
		{"\\x01-\\x02\\x03-\\x0a\\x1a-\\x1f", Range{2, 31}, "1", 1},
		{"\\x01-\\x02\\x03-\\x0a\\x1a-\\x1f", Range{4, 28}, "1, 3, 28-31", 5},
		{"\\x01-\\x02\\x03-\\x0a\\x1a-\\x1f", Range{4, 31}, "1, 3", 2},
	}
	prefix := trace.CallerName(0)

	for i, v := range testdata {
		c := &Charset{}
		c.MakeFromBytes([]byte(v.c))
		c.DifferenceRange(&v.r)
		str := c.StringAsInt()

		if str != v.str {
			t.Errorf("%s[%d] failed: str = %s, wanted = %s\n", prefix, i, str, v.str)
		}

		if c.Size() != v.size {
			t.Errorf("%s[%d] failed: size = %d, wanted = %d\n", prefix, i, c.Size(), v.size)
		}
	}
}

func TestCharsetDifferenceChar(t *testing.T) {
	testdata := []struct {
		ranges []Range
		ch     int32
		str    string
		size   uint32
	}{
		{[]Range{}, 1, "", 0},
		{[]Range{{1, 3}}, 0, "1-3", 2},
		{[]Range{{1, 3}}, 3, "1-3", 2},
		{[]Range{{1, 3}, {4, 10}}, 7, "1-3, 4-7, 8-10", 7},
		{[]Range{{1, 3}, {4, 10}}, 4, "1-3, 5-10", 7},
		{[]Range{{1, 3}, {7, 10}}, 5, "1-3, 7-10", 5},
	}
	prefix := trace.CallerName(0)

	for i, v := range testdata {
		c := &Charset{}
		c.UniteRangeSlice(v.ranges)
		c.DifferenceChar(v.ch)
		str := c.StringAsInt()

		if str != v.str {
			t.Errorf("%s[%d] failed: str = %s, wanted = %s\n", prefix, i, str, v.str)
		}

		if c.Size() != v.size {
			t.Errorf("%s[%d] failed: size = %d, wanted = %d\n", prefix, i, c.Size(), v.size)
		}
	}
}

func TestCharsetDifferenceCharset(t *testing.T) {
	testdata := []struct {
		c1   string
		c2   string
		str  string
		size uint32
	}{
		{"\\x01-\\x02\\x03-\\x0a\\x1a-\\x1f", "\\x00\\x04-\\x12", "1, 3, 26-31", 7},
		{"\\x01-\\x02\\x03-\\x0a\\x1a-\\x1f", "\\x00\\x04-\\x1c", "1, 3, 28-31", 5},
		{"\\x01-\\x02\\x03-\\x0a\\x1a-\\x1f", "\\x00\\x04-\\x2c", "1, 3", 2},
	}
	prefix := trace.CallerName(0)

	for i, v := range testdata {
		c1 := &Charset{}
		c1.MakeFromBytes([]byte(v.c1))
		c2 := &Charset{}
		c2.MakeFromBytes([]byte(v.c2))

		c1.DifferenceCharset(c2)
		str := c1.StringAsInt()

		if str != v.str {
			t.Errorf("%s[%d] failed: str = %s, wanted = %s\n", prefix, i, str, v.str)
		}

		if c1.Size() != v.size {
			t.Errorf("%s[%d] failed: size = %d, wanted = %d\n", prefix, i, c1.Size(), v.size)
		}
	}
}
