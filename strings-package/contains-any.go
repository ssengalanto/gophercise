package main

import (
	"fmt"
	"strings"
)

/*
	ContainsAny reports whether any Unicode code points in chars are within s.

	func ContainsAny(s, chars string) bool
*/

func main() {
	a := "Golang is cool"
	b := "g"
	c := "G"
	d := "z"

	fmt.Println(strings.ContainsAny(a,b)) // true
	fmt.Println(strings.ContainsAny(a,c)) // true
	fmt.Println(strings.ContainsAny(b,d)) // false
}
