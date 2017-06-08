package main

import (
//"charset"
//"fmt"
//"io"
//"os"
//"reflect"
)

type A struct {
	x int
}

func (a *A) f(int) bool {
	return false
}

func main() {

	/*
		var gen charset.CharsetGenForCpp

		gen.GenerateMask(os.Stdout, nil)

		var r1 charset.Range

		//fmt.Printf("0x%x\n", uint32(-1))

		r1 = charset.Range{1, 2}

		r1.PrintAsChar(os.Stdout).WriteString("\n")

		r1 = charset.Range{1, 6}
		r1.PrintAsChar(os.Stdout).WriteString("\n")

		r1 = charset.Range{0, 257}
		r1.Print(os.Stdout).WriteString("\n")
		r1.PrintEachChar(os.Stdout).WriteString("\n")

		fmt.Println("r1 = ", r1)
		fmt.Printf("%c\n", '\\')

		var a A

		p := a.f

		fmt.Println("type =", reflect.TypeOf(p))
	*/
}
