package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
	TrimFunc returns a slice of the string s with all leading and trailing Unicode code points c satisfying f(c) removed.
	func TrimFunc(s string, f func(rune) bool) string
*/

func main() {
	fmt.Print(strings.TrimFunc("--_Golang is cool!_--", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
}