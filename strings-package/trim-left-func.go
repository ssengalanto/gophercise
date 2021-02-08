package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
	TrimLeftFunc returns a slice of the string s with all leading Unicode code points c satisfying f(c) removed.
	func TrimLeftFunc(s string, f func(rune) bool) stringx
*/

func main() {
	fmt.Print(strings.TrimLeftFunc("--_Golang is cool!_--", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
}