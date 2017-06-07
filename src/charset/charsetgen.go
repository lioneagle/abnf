package charset

import (
	"basic"
	"container/list"
	"fmt"
	//"os"
	//"errors"
	//"strconv"
)

type charsetNode struct {
	Name       []byte
	MaskName   []byte
	MaskValue  int
	ActionName []byte
	VarIndex   int
	Set        Charset
}

type varData struct {
	TypeSize int
	TypeName []byte
	Data     [256]byte
}

type charsetNodeList struct {
	list.List
	MaskNameMaxLen   int
	ActionNameMaxLen int
	VarName          []byte
	VarList          []varData
}

func (charsets *charsetNodeList) Empty() bool {
	return charsets.Len() == 0
}

type CharsetGen struct {
	GenOneVar bool
	Usebit    bool
	VarName   string
}

type CharsetGenForC struct {
	CharsetGen

	VarTypeSize       int
	AutoSelectVarSize bool
}

func (gen *CharsetGenForC) GenerateMask(w basic.AbnfWriter, charsets *charsetNodeList) {
	fmt.Println("generator mask for C")
	if charsets == nil || charsets.Empty() || !gen.Usebit {
		return
	}

	MaxLen := charsets.MaskNameMaxLen + 4

	for e := charsets.Front(); e != nil; e = e.Next() {
		val := e.Value.(charsetNode)
		w.WriteString(fmt.Sprintf("#define %s", val.MaskName))
		basic.PrintIndent(w, MaxLen-len(val.MaskName))
		w.WriteString(fmt.Sprintf("(0x%08x)\r\n", val.MaskValue))
	}

}

func (gen *CharsetGenForC) GenerateAction(w basic.AbnfWriter, charsets *charsetNodeList) {
	fmt.Println("generator action for C")
	if charsets == nil || charsets.Empty() {
		return
	}

	MaxLen := charsets.ActionNameMaxLen + 4

	for e := charsets.Front(); e != nil; e = e.Next() {
		val := e.Value.(charsetNode)
		w.WriteString(fmt.Sprintf("#define %s", val.ActionName))
		basic.PrintIndent(w, MaxLen-len(val.ActionName))

		if gen.GenOneVar {

			if val.VarIndex == 0 {
				w.WriteString(fmt.Sprintf("(%s[(unsigned char)(ch)]", charsets.VarName))
			} else {
				w.WriteString(fmt.Sprintf("(%s[(unsigned char)(ch) + %d]", charsets.VarName, val.VarIndex*256))
			}
		} else if len(charsets.VarList) == 1 {
			w.WriteString(fmt.Sprintf("(%s[(unsigned char)(ch)]", charsets.VarName))
		} else {
			w.WriteString(fmt.Sprintf("(%s%d[(unsigned char)(ch)]", charsets.VarName, val.VarIndex))
		}

		if gen.Usebit {
			w.WriteString(fmt.Sprintf(" & %s", val.MaskName))
		}
		w.WriteString(")\r\n")
	}

}

func (gen *CharsetGenForC) GenerateVarDefinition(w basic.AbnfWriter, charsets *charsetNodeList) {
	fmt.Println("generator var defination for C")
}

func (gen *CharsetGenForC) GenerateVarDeclaration(w basic.AbnfWriter, charsets *charsetNodeList) {
	fmt.Println("generator var declaration for C")
}

type CharsetGenForCpp struct {
	CharsetGenForC

	ActionUseMacro bool
}

func (gen *CharsetGenForCpp) GenerateMask(w basic.AbnfWriter, charsets *charsetNodeList) {
	gen.CharsetGenForC.GenerateMask(w, charsets)
	fmt.Println("generator mask for Cpp")
}

func (gen *CharsetGenForCpp) GenerateAction(w basic.AbnfWriter, charsets *charsetNodeList) {
	fmt.Println("generator action for Cpp")
}

func (gen *CharsetGenForCpp) GenerateVarDefinition(w basic.AbnfWriter, charsets *charsetNodeList) {
	fmt.Println("generator var defination for Cpp")
}

func (gen *CharsetGenForCpp) GenerateVarDeclaration(w basic.AbnfWriter, charsets *charsetNodeList) {
	fmt.Println("generator var declaration for Cpp")
}

type CharsetGenForGolang struct {
	CharsetGenForC
}

func (gen *CharsetGenForGolang) GenerateMask(w basic.AbnfWriter, charsets *charsetNodeList) {
	fmt.Println("generator mask for C")
}

func (gen *CharsetGenForGolang) GenerateAction(w basic.AbnfWriter, charsets *charsetNodeList) {
	fmt.Println("generator action for C")
}

func (gen *CharsetGenForGolang) GenerateVarDefinition(w basic.AbnfWriter, charsets *charsetNodeList) {
	fmt.Println("generator var defination for C")
}

func (gen *CharsetGenForGolang) GenerateVarDeclaration(w basic.AbnfWriter, charsets *charsetNodeList) {
	fmt.Println("generator var declaration for C")
}

func (charsets *charsetNodeList) Calculate() {

}

type CharsetGenInterface interface {
	GenerateMask(w basic.AbnfWriter, charsets charsetNodeList)
	GenerateAction(w basic.AbnfWriter, charsets charsetNodeList)
	GenerateVarDefinition(w basic.AbnfWriter, charsets charsetNodeList)
	GenerateVarDeclaration(w basic.AbnfWriter, charsets charsetNodeList)
}
