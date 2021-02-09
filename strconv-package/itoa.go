package main

import (
	"fmt"
	"strconv"
)

/*
	Itoa is equivalent to FormatInt(int64(i), 10).
	func Itoa(i int) string
*/

func main() {
	i := 10
	s := strconv.Itoa(i)
	fmt.Printf("%T, %v\n", s, s)
}