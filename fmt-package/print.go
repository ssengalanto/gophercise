package main

import "fmt"

/*
Print formats using the default formats for its operands and writes to standard output.
Spaces are added between operands when neither is a string.
It returns the number of bytes written and any write error encountered.

func Print(a ...interface{}) (n int, err error)
*/

func main() {
	const name, age = "Hunter", 4
	p, err := fmt.Print(name, "is", age, "years old.")
	fmt.Println(p, err)
}
