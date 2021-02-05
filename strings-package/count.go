package main

import (
	"fmt"
	"strings"
)

/*
	Count counts the number of non-overlapping instances of substr in s.
	If substr is an empty string, Count returns 1 + the number of Unicode code points in s.

	func Count(s, substr string) int
*/

func main() {
	a := "Golang is cool"
	b := "o"
	c := "oo"
	d := "z"

	fmt.Println(strings.Count(a,b)) // 3
	fmt.Println(strings.Count(a,c)) // 1
	fmt.Println(strings.Count(a,d)) // 0
}
