package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.QuoteRuneToASCII('☺')
	fmt.Println(s)

	r := strconv.QuoteRuneToASCII('a')
	fmt.Println(r)
}
