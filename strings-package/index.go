package main

import (
	"fmt"
	"strings"
)

/*
	Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.

	func Index(s, substr string) int
*/

func main() {

	a := "Golang"

	fmt.Println(strings.Index(a, "Go"))
	fmt.Println(strings.Index(a, "lang"))
	fmt.Println(strings.Index(a, "go"))
}