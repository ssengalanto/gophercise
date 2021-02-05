package main

import (
	"fmt"
	"strings"
)

/*
	HasPrefix tests whether the string s begins with prefix.

	func HasPrefix(s, prefix string) bool
*/

func main() {
	a := "Golang"

	fmt.Println(strings.HasPrefix(a,"Go")) // true
	fmt.Println(strings.HasPrefix(a,"GO")) // false
	fmt.Println(strings.HasPrefix(a,"n/a")) // false
}
