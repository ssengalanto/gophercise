package main

import (
	"fmt"
	"strings"
)

/*
	TrimPrefix returns s without the provided leading prefix string. If s doesn't start with prefix, s is returned unchanged.

	func TrimPrefix(s, prefix string) string
*/

func main() {
	var s = "--_Golang is cool!_--"
	a := strings.TrimPrefix(s, "--_Golang ")
	b := strings.TrimPrefix(s, "--_golang, ")
	fmt.Println(a)
	fmt.Println(b)
}