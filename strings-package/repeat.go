package main

import (
	"fmt"
	"strings"
)

/*
	Repeat returns a new string consisting of count copies of the string s.
	It panics if count is negative or if the result of (len(s) * count) overflows.

	func Repeat(s string, count int) string
*/

func main() {
	fmt.Println(strings.Repeat("go ", 2) + "lang")
}