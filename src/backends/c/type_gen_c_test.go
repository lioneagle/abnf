package c

import (
	"fmt"
	"os"
	"testing"

	//"github.com/lioneagle/goutil/src/test"
	"github.com/lioneagle/abnf/src/gen/peg_gen"
)

func TestTypeInfoCalcPadNumber(t *testing.T) {
	struct1 := peg_gen.NewTypeInfoStruct("struct_A")
	byte1 := peg_gen.NewTypeInfoInt("byte")
	byte1.SetSize(1)

	int1 := peg_gen.NewTypeInfoInt("DWORD")
	int1.SetSize(4)

	char1 := peg_gen.NewTypeInfoChar("")

	string1 := peg_gen.NewTypeInfoString("")
	string1.SetSize(64)

	struct1.AppendMember(byte1, "x", "coordinate X")
	struct1.AppendMember(int1, "y", "coordinate Y")
	struct1.AppendMember(string1, "abName", "my name")

	struct2 := peg_gen.NewTypeInfoStruct("struct_B")
	int2 := peg_gen.NewTypeInfoInt("")
	int2.SetSize(8)

	int3 := peg_gen.NewTypeInfoInt("SHORT")
	int3.SetSize(2)

	struct2.AppendMember(int2, "s", "")
	struct2.AppendMember(int3, "t", "")

	struct1.AppendMember(struct2, "b", "")
	struct1.AppendMember(char1, "ch", "")

	struct1.CalcAlignSize()
	struct1.CalcPadNumber(8)

	generator := NewTypeGeneratorForC(peg_gen.NewConfig())

	generator.GenerateStruct(struct1, os.Stdout)
	fmt.Printf("\r\n")

	generator.GenerateStruct(struct2, os.Stdout)
	fmt.Printf("\r\n")

}
