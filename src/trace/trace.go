package trace

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

func Trace() *Stack {
	return TraceN(1, 32)
}

type Record struct {
	FuncName string
	File     string
	Line     int
}

type Stack struct {
	shortFileName      bool
	showFileName       bool
	showLineNo         bool
	shortFileNameDepth int
	indent             int
	data               []*Record
}

func NewStack(depth int) *Stack {
	return &Stack{
		shortFileName:      true,
		showFileName:       true,
		showLineNo:         true,
		shortFileNameDepth: 1,
		indent:             0,
		data:               make([]*Record, 0, depth),
	}
}

func (this *Stack) SetIndent(indent int) {
	this.indent = indent
}

func (this *Stack) SetShortFileNameDepth(depth int) {
	this.shortFileNameDepth = depth
}

func (this *Stack) EnableShortFileName() {
	this.shortFileName = true
}

func (this *Stack) DisableShortFileName() {
	this.shortFileName = false
}

func (this *Stack) ShowFileName() {
	this.showFileName = true
}

func (this *Stack) HideFileName() {
	this.showFileName = false
}

func (this *Stack) ShowLineNo() {
	this.showLineNo = true
}

func (this *Stack) HideLineNo() {
	this.showLineNo = false
}

func (this *Stack) String() string {
	var buf bytes.Buffer
	indentStr := fmt.Sprintf(fmt.Sprintf("%%%ds", this.indent), "")

	for i, v := range this.data {
		buf.WriteString(indentStr)

		filename := v.File
		if this.shortFileName {
			filename = extractFileName(filename, this.shortFileNameDepth)
		}

		fmt.Fprintf(&buf, "[%3d]: ", len(this.data)-i-1)

		if this.showFileName {
			fmt.Fprintf(&buf, "%s: ", filename)
		}
		if this.showLineNo {
			fmt.Fprintf(&buf, "%d: ", v.Line)
		}
		fmt.Fprintf(&buf, "%s\r\n", v.FuncName)
	}

	if len(this.data) != 0 {
		buf.WriteString(indentStr)
		buf.WriteString("... ...\r\n")
	}
	return buf.String()
}

func TraceN(skip, depth int) *Stack {
	stack := NewStack(depth)
	for i := 0; i < depth; i++ {
		v := Caller(skip + i + 1)
		if v == nil {
			break
		}
		stack.data = append(stack.data, v)
	}
	return stack
}

func Caller(skip int) *Record {
	pc, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return nil
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil || strings.HasPrefix(fn.Name(), "runtime.") {
		return nil
	}
	return &Record{
		FuncName: fn.Name(),
		File:     file,
		Line:     line,
	}
}

func CallerName(skip int) string {
	pc, _, _, ok := runtime.Caller(skip + 1)
	if !ok {
		return ""
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil || strings.HasPrefix(fn.Name(), "runtime.") {
		return ""
	}

	return fn.Name()
}

func extractFileName(fileName string, shortFileNameDepth int) string {
	if shortFileNameDepth < 1 {
		shortFileNameDepth = 1
	}

	i := 0
	depth := 0
	for i = len(fileName) - 1; i >= 0; i-- {
		if fileName[i] == '/' {
			depth++
			if depth >= shortFileNameDepth {
				break
			}
		}
	}
	return fileName[i+1 : len(fileName)]
}
