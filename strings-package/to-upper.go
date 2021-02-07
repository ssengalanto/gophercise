package main

import (
	"fmt"
	"strings"
)

/*
	ToUpper returns s with all Unicode letters mapped to their upper case.
	func ToUpper(s string) string
*/

func main() {
	fmt.Println(strings.ToUpper("Gopher"))
}