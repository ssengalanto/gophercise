package main

import (
	"fmt"
	"strings"
)

/*
	LastIndexAny returns the index of the last instance of any Unicode code point from chars in s,
	or -1 if no Unicode code point from chars is present in s.

	func LastIndexAny(s, chars string) int
*/

func main() {
	a := "go gopher"
	fmt.Println(strings.LastIndexAny(a, "go")) // 4
	fmt.Println(strings.LastIndexAny(a, "rodent")) // 8
	fmt.Println(strings.LastIndexAny(a, "fail")) // -1
}
