package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
	FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c) and returns an array of slices of s.
	If all code points in s satisfy f(c) or the string is empty, an empty slice is returned.
	FieldsFunc makes no guarantees about the order in which it calls f(c) and assumes that f always returns the same value for a given c.

	func FieldsFunc(s string, f func(rune) bool) []string
*/

func main() {
	f := func(c rune) bool {
		fmt.Println(c, string(c), unicode.IsLetter(c), unicode.IsNumber(c))
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", strings.FieldsFunc("  foo1;bar2,baz3...", f))
}