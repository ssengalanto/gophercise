package main

import (
	"fmt"
	"strings"
)

/*
	LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.

	func LastIndexByte(s string, c byte) int
*/

func main() {
	a := "golang is cool"
	fmt.Println(strings.LastIndexByte(a, 'l')) // 13
	fmt.Println(strings.LastIndexByte(a, 'o')) // 12
	fmt.Println(strings.LastIndexByte(a, 'x')) // -1
}