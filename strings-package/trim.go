package main

import (
	"fmt"
	"strings"
)

/*
	Trim returns a slice of the string s with all leading and trailing Unicode code points contained in cutset removed.
	func Trim(s, cutset string) string
*/

func main() {
	fmt.Print(strings.Trim("--_Golang is cool!_--", "-_"))
}