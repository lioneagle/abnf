package trace

import (
	//"fmt"
	"testing"
)

func TestCallerName(t *testing.T) {
	name := CallerName(0)
	if name != "trace.TestCallerName" {
		t.Errorf("TestCallerName failed: name = \"%s\", wanted = \"trace.TestCallerName\"\r\n", name)
	}

	name2 := CallerName(1)
	if name2 != "testing.tRunner" {
		t.Errorf("TestCallerName failed: name = \"%s\", wanted = \"testing.tRunner\"\r\n", name2)
	}
}

func TestTraceN(t *testing.T) {
	stack := TraceN(0, 2)
	stack.SetIndent(2)
	stack.EnableShortFileName()
	stack.SetShortFileNameDepth(1)
	stack.HideFileName()
	stack.HideLineNo()

	str := stack.String()
	wanted := "  [  1]: trace.TestTraceN\r\n  [  0]: testing.tRunner\r\n  ... ...\r\n"
	//fmt.Printf("str =\r\n%v\r\n", []byte(str))
	//fmt.Printf("wanted =\r\n%v\r\n", []byte(wanted))

	if str != wanted {
		t.Errorf("TestCallerName failed: str = \r\n%s, \r\nwanted = \r\n%s", str, wanted)
	}

}
