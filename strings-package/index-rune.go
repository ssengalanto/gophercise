package main

import (
	"fmt"
	"strings"
)

/*
	IndexRune returns the index of the first instance of the Unicode code point r, or -1 if rune is not present in s.
	If r is utf8.RuneError, it returns the first instance of any invalid UTF-8 byte sequence.

	func IndexRune(s string, r rune) int
*/

func main() {
	a := "Golang"
	fmt.Println(strings.IndexRune(a, 'g'))
	fmt.Println(strings.IndexRune(a, 'G'))
	fmt.Println(strings.IndexRune(a, 'z'))
}