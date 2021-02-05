package main

import (
	"fmt"
	"strings"
)

/*
	HasSuffix tests whether the string s ends with suffix.

	func HasSuffix(s, suffix string) bool
*/

func main() {
	a := "Golang"

	fmt.Println(strings.HasSuffix(a,"lang")) // true
	fmt.Println(strings.HasSuffix(a,"LANG")) // false
	fmt.Println(strings.HasSuffix(a,"n/a")) // false
}
