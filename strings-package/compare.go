package main

import (
	"fmt"
	"strings"
)

/*
	Compare returns an integer comparing two strings lexicographically. The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
	Compare is included only for symmetry with package bytes. It is usually clearer and always faster to use the built-in string comparison operators ==, <, >, and so on.

	func Compare(a, b string) int
*/

func main() {
	fmt.Println(strings.Compare("a", "b")) // -1
	fmt.Println(strings.Compare("a", "a")) // 0
	fmt.Println(strings.Compare("b", "a")) // 1
}
