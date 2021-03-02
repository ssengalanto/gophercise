package main

import (
	"fmt"
	"strconv"
)

func main() {
	if s, err := strconv.ParseBool("true"); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	if s, err := strconv.ParseBool("f"); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	if s, err := strconv.ParseBool("0"); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	if s, err := strconv.ParseBool("1"); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
}
