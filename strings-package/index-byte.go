package main

import (
	"fmt"
	"strings"
)

/*
	IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.

	func IndexByte(s string, c byte) int
*/

func main() {
	a := "Golang"

	fmt.Println(strings.IndexByte(a, 'G'))
	fmt.Println(strings.IndexByte(a, 'g'))
	fmt.Println(strings.IndexByte(a, 'z'))
}