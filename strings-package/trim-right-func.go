package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
	TrimRightFunc returns a slice of the string s with all trailing Unicode code points c satisfying f(c) removed.

	func TrimRightFunc(s string, f func(rune) bool) string
*/

func main() {
	fmt.Print(strings.TrimRightFunc("--_Golang is cool!_--", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
}