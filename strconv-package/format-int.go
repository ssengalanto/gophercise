package main

import (
	"fmt"
	"strconv"
)

/*
	FormatInt returns the string representation of i in the given base, for 2 <= base <= 36.
	The result uses the lower-case letters 'a' to 'z' for digit values >= 10.

	func FormatInt(i int64, base int) string
*/

func main() {
	v := int64(255)

	a := strconv.FormatInt(v, 2)
	fmt.Printf("%T, %v\n", a, a)

	b := strconv.FormatInt(v, 8)
	fmt.Printf("%T, %v\n", b, b)

	c := strconv.FormatInt(v, 10)
	fmt.Printf("%T, %v\n", c, c)

	d := strconv.FormatInt(v, 16)
	fmt.Printf("%T, %v\n", d, d)
}