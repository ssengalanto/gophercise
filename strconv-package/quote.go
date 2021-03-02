package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "Go lang is cool"
	a := strconv.Quote(s)
	fmt.Println(s)
	fmt.Println(a)
}
