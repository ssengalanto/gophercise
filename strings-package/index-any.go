package main

import (
	"fmt"
	"strings"
)

/*
	IndexAny returns the index of the first instance of any Unicode code point from chars in s,
	or -1 if no Unicode code point from chars is present in s.

	func IndexAny(s, chars string) int
*/

func main() {
	a := "Golang"

	fmt.Println(strings.IndexAny(a, "aeiou"))
	fmt.Println(strings.IndexAny(a, "uoiea"))
	fmt.Println(strings.IndexAny(a, "---z"))
}