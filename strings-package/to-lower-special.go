package main

import (
	"fmt"
	"strings"
	"unicode"
)

 /*
 	ToLowerSpecial returns a copy of the string s with all Unicode letters mapped to their lower case using the case mapping specified by c.
 	func ToLowerSpecial(c unicode.SpecialCase, s string) string
 */

func main() {
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))
}