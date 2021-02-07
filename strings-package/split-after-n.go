package main

import (
	"fmt"
	"strings"
)

/*
	SplitAfterN slices s into substrings after each instance of sep and returns a slice of those substrings.

	The count determines the number of substrings to return:
	n > 0: at most n substrings; the last substring will be the unsplit remainder.
	n == 0: the result is nil (zero substrings)
	n < 0: all substrings

	func SplitAfterN(s, sep string, n int) []string
*/

func main() {
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c,d,e", ",", 3))
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c,d,e", ",", 5))
}