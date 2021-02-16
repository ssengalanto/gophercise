package main

import (
	"fmt"
	"strconv"
)

/*
	FormatBool returns "true" or "false" according to the value of b.
	func FormatBool(b bool) string
*/

func main() {
	a := strconv.FormatBool(true)
	fmt.Printf("%T, %v\n", a, a)

	b := strconv.FormatBool(false)
	fmt.Printf("%T, %v\n", b, b)

}