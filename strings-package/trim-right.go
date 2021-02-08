package main

import (
	"fmt"
	"strings"
)

/*
	TrimRight returns a slice of the string s, with all trailing Unicode code points contained in cutset removed.
	To remove a suffix, use TrimSuffix instead.

	func TrimRight(s, cutset string) string
*/

func main() {
	fmt.Print(strings.TrimRight("--_Golang is cool!_--", "-_"))
}