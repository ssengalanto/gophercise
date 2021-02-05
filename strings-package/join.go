package main

import (
	"fmt"
	"strings"
)

/*
	Join concatenates the elements of its first argument to create a single string.
	The separator string sep is placed between elements in the resulting string.

	func Join(elems []string, sep string) string
*/

func main() {
	s := []string{"Golang", "is", "cool"}
	fmt.Println(strings.Join(s, " "))
	fmt.Println(strings.Join(s, ", "))
	fmt.Println(strings.Join(s, "-"))
}