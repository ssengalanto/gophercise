package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := []byte("Go lang is cool: ")
	b := '🚀'
	a = strconv.AppendQuoteRuneToGraphic(a, b)
	fmt.Println(string(a))
	fmt.Println(string(b))
}