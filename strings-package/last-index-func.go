package main

import (
	"fmt"
	"strings"
	"unicode"
)

 /*
 	LastIndexFunc returns the index into s of the last Unicode code point satisfying f(c), or -1 if none do.

 	func LastIndexFunc(s string, f func(rune) bool) int
 */

func main() {
	fmt.Println(strings.LastIndexFunc("go 123", unicode.IsNumber))
	fmt.Println(strings.LastIndexFunc("123 go", unicode.IsNumber))
	fmt.Println(strings.LastIndexFunc("go", unicode.IsNumber))
}