package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := []byte("append quote to graphic: ")
	b := "Go lang is cool 🚀"
	a = strconv.AppendQuoteToGraphic(a, b)
	fmt.Println(string(a))
	fmt.Println(b)
}