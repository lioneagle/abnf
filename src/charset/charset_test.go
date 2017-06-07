package charset

import (
//"fmt"
//"os"
//"strconv"
//"testing"
)

/*
func checkCharset(t *testing.T, name string, charset *Charset, ranges []Range) {
	var size int32 = 0

	for i := 0; i < len(ranges); i++ {
		size += ranges[i].High - ranges[i].Low
	}

	if size != charset.Size() {
		charset.Print(os.Stdout).WriteString("\n")
		t.Errorf("%s check failed, size = %d, wanted = %d, \nranges = %v", name, charset.Size(), size, ranges)
	}

	for i := 0; i < len(ranges); i++ {
		low := ranges[i].Low
		high := ranges[i].High
		for j := low; j < high; j++ {
			if !charset.Contains(j) {
				t.Errorf("%s check failed, %d is not contained in charset, ranges = %v", name, j, ranges)
			}
		}
	}
}

func TestCharsetUniteRange(t *testing.T) {
	testdata := []struct {
		r     Range
		check []Range
		str   string
	}{
		{Range{1, 2}, []Range{{1, 2}}, "1"},
		{Range{1, 2}, []Range{{1, 2}}, "1"},
		{Range{7, 9}, []Range{{1, 2}, {7, 9}}, "1, 7-9"},
		{Range{3, 4}, []Range{{1, 2}, {3, 4}, {7, 9}}, "1, 3, 7-9"},
		{Range{6, 10}, []Range{{1, 2}, {3, 4}, {6, 10}}, "1, 3, 6-10"},
		{Range{6, 9}, []Range{{1, 2}, {3, 4}, {6, 10}}, "1, 3, 6-10"},
		{Range{5, 8}, []Range{{1, 2}, {3, 4}, {5, 10}}, "1, 3, 5-10"},
		{Range{15, 18}, []Range{{1, 2}, {3, 4}, {5, 10}, {15, 18}}, "1, 3, 5-10, 15-18"},
		{Range{20, 22}, []Range{{1, 2}, {3, 4}, {5, 10}, {15, 18}, {20, 22}}, "1, 3, 5-10, 15-18, 20-22"},
		{Range{7, 19}, []Range{{1, 2}, {3, 4}, {5, 19}, {20, 22}}, "1, 3, 5-19, 20-22"},
		{Range{4, 21}, []Range{{1, 2}, {3, 22}}, "1, 3-22"},
	}

	prefix := FuncName()
	var charset Charset

	for i, v := range testdata {
		name := fmt.Sprintf("%s[%d]", prefix, i)
		charset.UniteRange(&v.r)
		checkPrintResult(t, name, charset.Print, v.str)
		checkCharset(t, name, &charset, v.check)
	}
}

func makeCharset(charset *Charset, ranges []Range) {
	charset.RemoveAll()

	for i := 0; i < len(ranges); i++ {
		charset.UniteRange(&ranges[i])
	}
}

func TestCharsetUniteChar(t *testing.T) {

	var charset Charset

	checkPrintResult(t, "Charset.Print", charset.Print, "")

	charset.UniteRange(&Range{1, 2})
	charset.UniteRange(&Range{3, 4})
	charset.UniteChar(2)
	checkPrintResult(t, "Charset.UniteChar", charset.Print, "1-4")
	checkCharset(t, "Charset.UniteChar", &charset, []Range{{1, 4}})
}

func TestCharsetContains(t *testing.T) {

	var charset Charset

	makeCharset(&charset, []Range{{1, 2}, {3, 4}})

	if !charset.Contains(3) {
		t.Errorf("Charset.Contains failed, 3 should be contained in charset")
	}

	if charset.Contains(2) {
		t.Errorf("Charset.Contains failed, 2 should not be contained in charset")
	}

	makeCharset(&charset, []Range{{1, 2}, {3, 5}, {7, 10}})

	if !charset.Contains(4) {
		t.Errorf("Charset.Contains failed, 3 should be contained in charset")
	}

	if charset.Contains(6) {
		t.Errorf("Charset.Contains failed, 2 should not be contained in charset")
	}
}

type testCharsetUniteCharsetItem struct {
	c1    []Range
	c2    []Range
	check []Range
	str   string
}

func TestCharsetUniteCharset(t *testing.T) {

	var charset1, charset2 Charset

	data := []testCharsetUniteCharsetItem{
		{[]Range{{1, 2}}, []Range{}, []Range{{1, 2}}, "1"},
		{[]Range{{1, 2}}, []Range{{1, 2}}, []Range{{1, 2}}, "1"},
		{[]Range{{1, 2}, {3, 4}, {6, 10}}, []Range{{5, 8}, {9, 12}, {15, 16}}, []Range{{1, 2}, {3, 4}, {5, 12}, {15, 16}}, "1, 3, 5-12, 15"},
	}

	for i := 0; i < len(data); i++ {
		makeCharset(&charset1, data[i].c1)
		makeCharset(&charset2, data[i].c2)
		name := "Charset.UniteCharset[" + strconv.Itoa(i) + "]"
		charset1.UniteCharset(&charset2)
		checkPrintResult(t, name, charset1.Print, data[i].str)
		checkCharset(t, name, &charset1, data[i].check)
	}

}

type testCharsetPrintItem struct {
	c           []Range
	str         string
	strAsChar   string
	strEachChar string
}

func TestCharsetPrint(t *testing.T) {

	var charset Charset

	data := []testCharsetPrintItem{
		{[]Range{}, "", "", ""},
		{[]Range{{1, 2}}, "1", "\\x01", "\\x01"},
		{[]Range{{1, 2}, {'a', 'd'}}, "1, 97-100", "\\x01, a-d", "\\x01, a, b, c"},
	}

	for i := 0; i < len(data); i++ {
		makeCharset(&charset, data[i].c)
		name := "CharsetPrint[" + strconv.Itoa(i) + "]"
		checkPrintResult(t, name, charset.Print, data[i].str)
		checkPrintResult(t, name, charset.PrintAsChar, data[i].strAsChar)
		checkPrintResult(t, name, charset.PrintEachChar, data[i].strEachChar)
	}
}

func TestCharsetMakeFromBytes(t *testing.T) {

	var charset Charset

	err := charset.MakeFromBytes([]byte(""))
	if err != nil {
		t.Errorf("Charset.MakeFromBytes failed, err = %v", err)
	}
	checkPrintResult(t, "Charset.MakeFromBytes", charset.PrintAsChar, "")

	err = charset.MakeFromBytes([]byte("\\x01\\002-\\005a-fAY\\p\\x0f-\\x1A\\x9\\a\\b\\r\\n\\f\\t\\v\\"))
	if err != nil {
		t.Errorf("Charset.MakeFromBytes failed, err = %v", err)
	}
	checkPrintResult(t, "Charset.MakeFromBytes", charset.PrintAsChar, "\\x01-\\x05, \\a-\\x0e, \\x0f-\\x1a, 9, A, Y, \\\\, a-f, p, x")

	err = charset.MakeFromBytes([]byte("\\1+\\x1"))
	if err != nil {
		t.Errorf("Charset.MakeFromBytes failed, err = %v", err)
	}
	checkPrintResult(t, "Charset.MakeFromBytes", charset.PrintAsChar, "+, 1, \\\\, x")

	err = charset.MakeFromBytes([]byte("\\21"))
	if err != nil {
		t.Errorf("Charset.MakeFromBytes failed, err = %v", err)
	}
	checkPrintResult(t, "Charset.MakeFromBytes", charset.PrintAsChar, "1-3, \\\\")

	err = charset.MakeFromBytes([]byte("\\x05-\\x01"))
	if err != nil {
		t.Errorf("Charset.MakeFromBytes failed, err = %v", err)
	}
	checkPrintResult(t, "Charset.MakeFromBytes", charset.PrintAsChar, "\\x01-\\x05")

}

func TestUnescapeCharForEmptyBytes(t *testing.T) {
	ch, pos := unescapeChar([]byte(""), 0)
	if ch >= 0 {
		t.Errorf("unescapeChar shuold return -1, ret = %d, pos = %d", ch, pos)
	}
}
*/
