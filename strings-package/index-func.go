package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
	IndexFunc returns the index into s of the first Unicode code point satisfying f(c), or -1 if none do.

	func IndexFunc(s string, f func(rune) bool) int
*/

func main() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f))
	fmt.Println(strings.IndexFunc("Hello, world", f))
}