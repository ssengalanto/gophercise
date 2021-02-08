package main

import (
	"fmt"
	"strings"
)

/*
	TrimLeft returns a slice of the string s with all leading Unicode code points contained in cutset removed.
	To remove a prefix, use TrimPrefix instead.
	func TrimLeft(s, cutset string) string
*/

func main() {
	fmt.Print(strings.TrimLeft("--_Golang is cool!_--", "_-"))
}