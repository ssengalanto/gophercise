package main

import (
	"fmt"
	"strings"
)

/*
	TrimSuffix returns s without the provided trailing suffix string. If s doesn't end with suffix, s is returned unchanged.

	func TrimSuffix(s, suffix string) string
*/

func main() {
	var s = "--_Golang is cool!_--"
	a := strings.TrimSuffix(s, "cool!_--")
	b := strings.TrimSuffix(s, "cool!")
	fmt.Println(a)
	fmt.Println(b)
}