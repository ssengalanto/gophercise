package main

import (
	"fmt"
	"os"
)

/*
	Fprint: formats using the default formats for its operands and writes to w.
	Spaces are added between operands when neither is a string.
	It returns the number of bytes written and any write error encountered.

	func Fprint(w io.Writer, a ...interface{}) (n int, err error)
*/

func main() {
	const name, age = "Hunter", 4
	n, err := fmt.Fprint(os.Stdout, name, " is ", age, " years old.\n")

	// The n and err return values from Fprint are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
	}
	fmt.Print(n, " bytes written.\n")
}
