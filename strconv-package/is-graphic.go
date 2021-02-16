package main

import (
	"fmt"
	"strconv"
)

/*
	IsGraphic reports whether the rune is defined as a Graphic by Unicode.
	Such characters include letters, marks, numbers, punctuation, symbols, and spaces, from categories L, M, N, P, S, and Zs.
	func IsGraphic(r rune) bool
*/

func main() {
	a := strconv.IsGraphic('a')
	fmt.Println(a)

	b := strconv.IsGraphic('🚀')
	fmt.Println(b)

	c := strconv.IsGraphic('☘')
	fmt.Println(c)

	d := strconv.IsGraphic('\007')
	fmt.Println(d)
}