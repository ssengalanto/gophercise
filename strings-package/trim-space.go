package main

import (
	"fmt"
	"strings"
)

/*
	TrimSpace returns a slice of the string s, with all leading and trailing white space removed, as defined by Unicode.
	func TrimSpace(s string) string
*/

func main() {
	fmt.Println(strings.TrimSpace(" \t\n Golang is cool \n\t\r\n"))
}