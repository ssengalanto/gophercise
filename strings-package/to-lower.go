package main

import (
	"fmt"
	"strings"
)

/*
	ToLower returns s with all Unicode letters mapped to their lower case.
	func ToLower(s string) string
*/

func main() {
	fmt.Println(strings.ToLower("Gopher"))
}