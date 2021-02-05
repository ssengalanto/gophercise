package main

import (
	"fmt"
	"strings"
)

/*
	Contains reports whether substr is within s.

	func Contains(s, substr string) bool
*/

func main() {
	a := "Golang is cool"
	b := "Golang"
	c := "Go"
	d := "asdf"

	fmt.Println(strings.Contains(a,b)) // true
	fmt.Println(strings.Contains(b,c)) // true
	fmt.Println(strings.Contains(a,d)) // false
}
