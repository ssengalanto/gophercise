package main

import (
	"fmt"
	"strings"
)

/*
	SplitN slices s into substrings separated by sep and returns a slice of the substrings between those separators.

	The count determines the number of substrings to return:
	n > 0: at most n substrings; the last substring will be the unsplit remainder.
	n == 0: the result is nil (zero substrings)
	n < 0: all substrings

	func SplitN(s, sep string, n int) []string
*/

func main() {
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
	z := strings.SplitN("a,b,c", ",", 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)
}