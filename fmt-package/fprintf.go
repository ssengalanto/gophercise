package main

import (
	"fmt"
	"os"
)

/*
	Fprintf formats according to a format specifier and writes to w.
	It returns the number of bytes written and any write error encountered.

	func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
*/

func main() {
	const name, age = "Hunter", 4
	n, err := fmt.Fprintf(os.Stdout, "%s is %d years old.\n", name, age)

	// The n and err return values from Fprintf are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
	}
	fmt.Printf("%d bytes written.\n", n)

}
