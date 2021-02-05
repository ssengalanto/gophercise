package main

import (
	"fmt"
	"strings"
)

/*
	CContainsRune reports whether the Unicode code point r is within s.

	func ContainsRune(s string, r rune) bool
*/

func main() {
	a := "Golang is cool"
	b := "TypeScript"
	// Finds whether a string contains a particular Unicode code point.
	// The code point for the lowercase letter "a", for example, is 97.
	fmt.Println(strings.ContainsRune(a, 97)) // true
	fmt.Println(strings.ContainsRune(b, 97)) // false
}