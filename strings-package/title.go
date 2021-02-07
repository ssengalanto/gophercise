package main

import (
	"fmt"
	"strings"
)

/*
	Title returns a copy of the string s with all Unicode letters that begin words mapped to their Unicode title case.
	BUG(rsc): The rule Title uses for word boundaries does not handle Unicode punctuation properly.

	func Title(s string) string
*/

func main() {
	fmt.Println(strings.Title("golang is cool"))
}