package main

import (
	"fmt"
	"os"
	"strings"
)

/*
	Fscanf scans text read from r, storing successive space-separated values into successive arguments as determined by the format.
	It returns the number of items successfully parsed. Newlines in the input must match newlines in the format.

	func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
*/

func main() {
	var (
		i int
		b bool
		s string
	)
	r := strings.NewReader("5 true gophers")
	n, err := fmt.Fscanf(r, "%d %t %s", &i, &b, &s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
	}
	fmt.Println(i, b, s)
	fmt.Println(n)
}
