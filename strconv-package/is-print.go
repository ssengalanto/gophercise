package main

import (
	"fmt"
	"strconv"
)

/*
	IsPrint reports whether the rune is defined as printable by Go, with the same definition as unicode.
	IsPrint: letters, numbers, punctuation, symbols and ASCII space.
	func IsPrint(r rune) bool
*/

func main() {
	a := strconv.IsPrint('a')
	fmt.Println(a)

	b := strconv.IsGraphic('🚀')
	fmt.Println(b)

	c := strconv.IsPrint('\007')
	fmt.Println(c)
}