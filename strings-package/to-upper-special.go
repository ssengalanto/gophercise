package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
	ToUpperSpecial returns a copy of the string s with all Unicode letters mapped to their upper case using the case mapping specified by c.
	func ToUpperSpecial(c unicode.SpecialCase, s string) string
*/

func main() {
	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))
}