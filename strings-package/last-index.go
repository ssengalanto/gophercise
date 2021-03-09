package main

import (
	"fmt"
	"strings"
)

/*
	LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.

	func LastIndex(s, substr string) int
*/

func main() {
	a := "go gopher"
	fmt.Println(strings.Index(a, "go")) // first match.go 0
	fmt.Println(strings.LastIndex(a, "go")) // last match.go 3
	fmt.Println(strings.LastIndex(a, "rodent")) // no match.go -1
}
