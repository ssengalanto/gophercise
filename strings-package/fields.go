package main

/*
	Fields splits the string s around each instance of one or more consecutive white space characters, as defined by unicode.
	IsSpace, returning a slice of substrings of s or an empty slice if s contains only white space.

	func Fields(s string) []string
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Fields are: %q", strings.Fields("  foo   bar baz  abc  "))
}