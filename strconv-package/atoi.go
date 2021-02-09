package main

import (
	"fmt"
	"strconv"
)

/*
	Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.
	func Atoi(s string) (int, error)
*/

func main() {
	v := "10"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v", s, s)
	}
}