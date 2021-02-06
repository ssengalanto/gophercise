package main

import (
	"fmt"
	"strings"
)

/*
	Map returns a copy of the string s with all its characters modified according to the mapping function.
	If mapping returns a negative value, the character is dropped from the string with no replacement.

	func Map(mapping func(rune) rune, s string) string
*/

func main() {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
}