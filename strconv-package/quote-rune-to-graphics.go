package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.QuoteRuneToGraphic('☺')
	fmt.Println(s)

	s = strconv.QuoteRuneToGraphic('\u263a')
	fmt.Println(s)

	s = strconv.QuoteRuneToGraphic('\U0001f680')
	fmt.Println(s)

	s = strconv.QuoteRuneToGraphic('	') // tab character
	fmt.Println(s)
}
