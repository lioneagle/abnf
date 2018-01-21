package charset

import (
	"fmt"
	"testing"

	"github.com/lioneagle/goutil/src/test"
)

func TestCharsetUniteRange(t *testing.T) {
	testdata := []struct {
		ranges []Range
		str    string
		size   uint32
	}{
		{[]Range{{0, 0}, {}}, "", 0},
		{[]Range{{0, 0}, {1, 2}}, "1", 1},
		{[]Range{{1, 2}, {7, 9}}, "1, 7-8", 3},
		{[]Range{{7, 9}, {1, 2}}, "1, 7-8", 3},
		{[]Range{{7, 9}, {11, 15}}, "7-8, 11-14", 6},
		{[]Range{{7, 10}, {2, 12}}, "2-11", 10},
		{[]Range{{7, 10}, {2, 7}}, "2-9", 8},
		{[]Range{{7, 10}, {10, 13}}, "7-12", 6},
		{[]Range{{7, 10}, {8, 12}}, "7-11", 5},
		{[]Range{{7, 10}, {1, 9}}, "1-9", 9},

		{[]Range{{1, 2}}, "1", 1},                                                                                        // empty
		{[]Range{{1, 2}, {1, 2}}, "1", 1},                                                                                // same node
		{[]Range{{1, 2}, {2, 4}}, "1-3", 3},                                                                              // concat
		{[]Range{{1, 2}, {-2, -1}, {-5, -3}, {3, 5}}, "-5--4, -2, 1, 3-4", 6},                                            // insert before
		{[]Range{{1, 2}, {7, 9}}, "1, 7-8", 3},                                                                           // append at tail without cross
		{[]Range{{1, 2}, {7, 9}, {3, 4}}, "1, 3, 7-8", 4},                                                                // insert at middle without cross
		{[]Range{{1, 2}, {7, 9}, {3, 4}, {6, 10}}, "1, 3, 6-9", 6},                                                       // insert at middle without cross
		{[]Range{{1, 2}, {7, 9}, {3, 4}, {6, 10}, {6, 9}}, "1, 3, 6-9", 6},                                               // contained in one node
		{[]Range{{1, 2}, {7, 9}, {3, 4}, {6, 10}, {6, 9}, {5, 8}}, "1, 3, 5-9", 7},                                       // low is in one node and high is greater than last node's high
		{[]Range{{1, 2}, {7, 9}, {3, 4}, {6, 10}, {6, 9}, {5, 8}, {15, 18}}, "1, 3, 5-9, 15-17", 10},                     // append at tail without cross
		{[]Range{{1, 2}, {7, 9}, {3, 4}, {6, 10}, {6, 9}, {5, 8}, {15, 18}, {20, 22}}, "1, 3, 5-9, 15-17, 20-21", 12},    // append at tail without cross
		{[]Range{{1, 2}, {7, 9}, {3, 4}, {6, 10}, {6, 9}, {5, 8}, {15, 18}, {20, 22}, {7, 19}}, "1, 3, 5-18, 20-21", 18}, // low is in one node and high is greater than other two node
		{[]Range{{1, 2}, {3, 22}}, "1, 3-21", 20},                                                                        // low is equal to one node's high
		{[]Range{{1, 2}, {3, 22}, {23, 25}}, "1, 3-21, 23-24", 22},                                                       // low is equal to one node's high
		{[]Range{{1, 2}, {3, 22}, {23, 25}, {22, 23}}, "1, 3-24", 23},                                                    // high is equal to one node's low
		{[]Range{{1, 2}, {3, 22}, {23, 25}, {22, 23}, {100, 102}}, "1, 3-24, 100-101", 25},                               // append at tail without cross
		{[]Range{{1, 2}, {3, 22}, {23, 25}, {22, 23}, {100, 102}, {40, 42}}, "1, 3-24, 40-41, 100-101", 27},              // low is less than one node's low and high is greater than onther node's high
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			c := NewCharset()
			c.UniteRangeSlice(v.ranges)
			str := c.StringAsInt()

			test.EXPECT_EQ(t, str, v.str, "")
			test.EXPECT_EQ(t, c.Size(), v.size, "")
		})
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

	c := NewCharset()
	c.UniteRangeSlice(ranges)

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			test.EXPECT_EQ(t, c.Contains(v.val), v.contains, "")
		})
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
		{[]Range{{1, 3}}, 0, "0-2", 3},
		{[]Range{{1, 3}}, 3, "1-3", 3},
		{[]Range{{1, 3}, {4, 10}}, 3, "1-9", 9},
		{[]Range{{1, 3}, {4, 10}}, 7, "1-2, 4-9", 8},
		{[]Range{{1, 3}, {7, 10}}, 5, "1-2, 5, 7-9", 6},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			c := NewCharset()
			c.UniteRangeSlice(v.ranges)
			c.UniteChar(v.ch)
			str := c.StringAsInt()

			test.EXPECT_EQ(t, str, v.str, "")
			test.EXPECT_EQ(t, c.Size(), v.size, "")
		})
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
		{"\\x01\\002", print_as_int, "1-2", 2},
		{"\\x01\\002-\\x05", print_as_int, "1-5", 5},
		{"a-", print_as_char, "\\-, a", 2},
		{"\\21", print_as_char, "1-2, \\\\", 3},
		{"\\x05-\\x01", print_as_char, "\\x01-\\x05", 5},
		{"a-d", print_each_char, "a, b, c, d", 4},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			c := NewCharset()
			c.MakeFromBytes([]byte(v.src))
			str := c.toString(v.printType)

			test.EXPECT_EQ(t, str, v.str, "")
			test.EXPECT_EQ(t, c.Size(), v.size, "")
		})
	}
}

func TestCharsetRemoveAll(t *testing.T) {
	c := NewCharset()
	c.MakeFromBytes([]byte("\\x01-\\x05a-c\\k-m"))
	c.RemoveAll()
	str := c.StringAsInt()

	test.EXPECT_EQ(t, str, "", "")
	test.EXPECT_EQ(t, c.Size(), uint32(0), "")
}

func TestCharsetDifferenceRange(t *testing.T) {
	testdata := []struct {
		c    string
		r    Range
		str  string
		size uint32
	}{
		{"", Range{1, 2}, "", 0},
		{"\\x01-\\x02", Range{1, 3}, "", 0},
		{"\\x03-\\x0a", Range{2, 6}, "6-10", 5},
		{"\\x03-\\x0a", Range{2, 11}, "", 0},
		{"\\x03-\\x0a", Range{3, 12}, "", 0},
		{"\\x03-\\x0a", Range{4, 11}, "3", 1},
		{"\\x03-\\x0a", Range{5, 8}, "3-4, 8-10", 5},
		{"\\x03-\\x0a", Range{1, 4}, "4-10", 7},
		{"\\x03-\\x0a", Range{10, 13}, "3-9", 7},

		{"\\x01-\\x02\\x0a-\\x0f\\x1a-\\x1f", Range{1, 3}, "10-15, 26-31", 12},
		{"\\x01-\\x02\\x04-\\x0a\\x1a-\\x1f", Range{2, 5}, "1, 5-10, 26-31", 13},
		{"\\x01-\\x02\\x04-\\x0a\\x1a-\\x1f", Range{2, 11}, "1, 26-31", 7},
		{"\\x01-\\x02\\x04-\\x0a\\x1a-\\x1f", Range{3, 12}, "1-2, 26-31", 8},
		{"\\x01-\\x02\\x04-\\x0a\\x1a-\\x1f", Range{5, 11}, "1-2, 4, 26-31", 9},
		{"\\x01-\\x02\\x04-\\x0a\\x1a-\\x1f", Range{5, 8}, "1-2, 4, 8-10, 26-31", 12},
		{"\\x01-\\x02\\x04-\\x0a\\x1a-\\x1f", Range{1, 4}, "4-10, 26-31", 13},
		{"\\x01-\\x02\\x04-\\x0a\\x1a-\\x1f", Range{10, 13}, "1-2, 4-9, 26-31", 14},

		{"\\x01-\\x02\\x04-\\x0a\\x1a-\\x1f", Range{2, 29}, "1, 29-31", 4},
		{"\\x01-\\x02\\x04-\\x0a\\x1a-\\x1f", Range{2, 32}, "1", 1},
		{"\\x01-\\x02\\x04-\\x0a\\x1a-\\x1f", Range{5, 29}, "1-2, 4, 29-31", 6},
		{"\\x01-\\x02\\x04-\\x0a\\x1a-\\x1f", Range{5, 32}, "1-2, 4", 3},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			c := NewCharset()
			c.MakeFromBytes([]byte(v.c))
			c.DifferenceRange(&v.r)
			str := c.StringAsInt()

			test.EXPECT_EQ(t, str, v.str, "")
			test.EXPECT_EQ(t, c.Size(), v.size, "")
		})
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
		{[]Range{{1, 3}}, 0, "1-2", 2},
		{[]Range{{1, 3}}, 3, "1-2", 2},
		{[]Range{{1, 3}, {4, 10}}, 7, "1-2, 4-6, 8-9", 7},
		{[]Range{{1, 3}, {4, 10}}, 4, "1-2, 5-9", 7},
		{[]Range{{1, 3}, {7, 10}}, 5, "1-2, 7-9", 5},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			c := NewCharset()
			c.UniteRangeSlice(v.ranges)
			c.DifferenceChar(v.ch)
			str := c.StringAsInt()

			test.EXPECT_EQ(t, str, v.str, "")
			test.EXPECT_EQ(t, c.Size(), v.size, "")
		})
	}
}

func TestCharsetDifferenceCharset(t *testing.T) {
	testdata := []struct {
		c1   string
		c2   string
		str  string
		size uint32
	}{
		{"\\x01-\\x02\\x04-\\x0a\\x1a-\\x1f", "\\x00\\x05-\\x12", "1-2, 4, 26-31", 9},
		{"\\x01-\\x02\\x04-\\x0a\\x1a-\\x1f", "\\x00\\x05-\\x1c", "1-2, 4, 29-31", 6},
		{"\\x01-\\x02\\x04-\\x0a\\x1a-\\x1f", "\\x00\\x05-\\x2c", "1-2, 4", 3},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			c1 := NewCharset()
			c1.MakeFromBytes([]byte(v.c1))
			c2 := NewCharset()
			c2.MakeFromBytes([]byte(v.c2))
			c1.DifferenceCharset(c2)
			str := c1.StringAsInt()

			test.EXPECT_EQ(t, str, v.str, "")
			test.EXPECT_EQ(t, c1.Size(), v.size, "")
		})
	}
}

func TestCharsetMakeFromBytesInverse(t *testing.T) {
	testdata := []struct {
		any       Range
		src       string
		printType int
		str       string
		size      uint32
	}{
		{Range{1, 256}, "", print_as_int, "1-255", 255},
		{Range{1, 10}, "\\x01\\002", print_as_int, "3-9", 7},
		{Range{1, 30}, "\\x01\\002-\\x05", print_as_int, "6-29", 24},
		{Range{'a', 'g'}, "a-d", print_each_char, "e, f", 2},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			c := NewCharset()
			c.MakeFromBytesInverse(&v.any, []byte(v.src))
			str := c.toString(v.printType)

			test.EXPECT_EQ(t, str, v.str, "")
			test.EXPECT_EQ(t, c.Size(), v.size, "")
		})
	}
}

func TestCharsetEqual(t *testing.T) {
	testdata := []struct {
		c1    string
		c2    string
		equal bool
	}{
		{"\\x01-\\x02\\x03-\\x0a\\x1a-\\x1f", "\\x01-\\x02\\x03-\\x0a\\x1a-\\x1f", true},
		{"a-e", "a-cde", true},
		{"a-c", "A-C", false},
		{"a-c", "8-10a-b", false},
	}

	for i, v := range testdata {
		v := v
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()

			c1 := NewCharset()
			c1.MakeFromBytes([]byte(v.c1))
			c2 := NewCharset()
			c2.MakeFromBytes([]byte(v.c2))

			test.EXPECT_EQ(t, c1.Equal(c2), v.equal, "")
		})
	}
}
