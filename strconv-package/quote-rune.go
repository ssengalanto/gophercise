package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := '🚀'
	r := '☺'
	a := strconv.QuoteRuneToASCII(s)
	b := strconv.QuoteRuneToASCII(r)
	fmt.Println(s, r)
	fmt.Println(a, b)
}
