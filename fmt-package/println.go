package main

import "fmt"

/*
Println formats using the default formats for its operands and writes to standard output.
Spaces are always added between operands and a newline is appended.
It returns the number of bytes written and any write error encountered.

func Println(a ...interface{}) (n int, err error)
*/

func main() {
	const name, age = "Hunter", 4
	p, err := fmt.Println(name, "is", age, "years old.")
	fmt.Println(p, err)
}
