package main

import (
	"fmt"
	"strings"
)

/*
	Replace returns a copy of the string s with the first n non-overlapping instances of old replaced by new.
	If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string.
	If n < 0, there is no limit on the number of replacements.

	func Replace(s, old, new string, n int) string
*/

func main() {
	fmt.Println(strings.Replace("typescript typescript typescript", "typescript", "golang", 2))
	fmt.Println(strings.Replace("nodejs nodejs nodejs", "nodejs", "golang", -1))
}