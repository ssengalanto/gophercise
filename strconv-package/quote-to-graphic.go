package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.QuoteToGraphic("☺")
	fmt.Println(s, "☺")

	s = strconv.QuoteToGraphic("This is a \u263a	\u000a")
	fmt.Println(s, "\u263a	\u000a")

	s = strconv.QuoteToGraphic(`" This is a ☺ \n "`)
	fmt.Println(s, " This is a ☺ \n ")

}
