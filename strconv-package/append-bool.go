package main

import (
	"fmt"
	"strconv"
)

/*
	AppendBool appends "true" or "false", according to the value of b, to dst and returns the extended buffer.
	func AppendBool(dst []byte, b bool) []byte
*/

func main() {
	b := []byte("Golang is cool:")
	b = strconv.AppendBool(b, true)
	fmt.Println(string(b))
}