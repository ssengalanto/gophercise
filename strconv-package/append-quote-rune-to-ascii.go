package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("quote (ascii):")
	b = strconv.AppendQuoteToASCII(b, `"Go lang is cool"`)
	fmt.Println(string(b))
	fmt.Println(b)
}