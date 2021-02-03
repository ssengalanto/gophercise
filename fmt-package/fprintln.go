package main

import (
	"fmt"
	"os"
)

/*
	Fprintln formats using the default formats for its operands and writes to w.
	Spaces are always added between operands and a newline is appended.
	It returns the number of bytes written and any write error encountered.

	func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
*/

func main() {
	const name, age = "Hunter", 4
	n, err := fmt.Fprintln(os.Stdout, name, "is", age, "years old.")

	// The n and err return values from Fprintln are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintln: %v\n", err)
	}
	fmt.Println(n, "bytes written.")

}