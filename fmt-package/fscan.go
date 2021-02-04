package main

import (
	"fmt"
	"os"
	"strings"
)

/*
	Fscan scans text read from r, storing successive space-separated values into successive arguments.
	Newlines count as space. It returns the number of items successfully scanned.
	If that is less than the number of arguments, err will report why.

	func Fscan(r io.Reader, a ...interface{}) (n int, err error)
*/

func main() {
	var (
		i       int
		b       bool
		s, t, u string
	)
	r := strings.NewReader(`5 true gophers
	new line`)
	n, err := fmt.Fscan(r, &i, &b, &s, &t, &u)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
	}
	fmt.Println(i, b, s, t, u)
	fmt.Println(n)
}
