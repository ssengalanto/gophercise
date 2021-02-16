package main

import (
	"fmt"
	"strconv"
)

/*
	FormatFloat converts the floating-point number f to a string, according to the format fmt and precision prec.
	It rounds the result assuming that the original was obtained from a floating-point value of bitSize bits (32 for float32, 64 for float64).

	func FormatFloat(f float64, fmt byte, prec, bitSize int) string
*/

func main() {
	v := 3.1415926535

	s32 := strconv.FormatFloat(v, 'E', -1, 32)
	fmt.Printf("%T, %v\n", s32, s32)

	s64 := strconv.FormatFloat(v, 'E', -1, 64)
	fmt.Printf("%T, %v\n", s64, s64)

}