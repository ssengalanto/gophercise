package main

import (
	"fmt"
	"strings"
)

/*
	EqualFold reports whether s and t, interpreted as UTF-8 strings, are equal under Unicode case-folding,
	which is a more general form of case-insensitivity.

	func EqualFold(s, t string) bool
*/

func main() {
	a := "Go"
	b := "go"
	c := "GO"

	fmt.Println(strings.EqualFold(a,b)) // true
	fmt.Println(strings.EqualFold(a,c)) // true
}
