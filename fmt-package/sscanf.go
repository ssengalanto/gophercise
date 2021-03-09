package main

import (
	"fmt"
)

/*
	Sscanf scans the argument string, storing successive space-separated values into
	successive arguments as determined by the format.
	It returns the number of items successfully parsed.
	Newlines in the input must match.go newlines in the format.

	func Sscanf(str string, format string, a ...interface{}) (n int, err error)
*/

func main() {
	var name string
	var age int
	n, err := fmt.Sscanf("Hunter is 4 years old", "%s is %d years old", &name, &age)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d: %s, %d\n", n, name, age)
}
