package main

import (
	"fmt"
)

/*
	Sscan scans the argument string, storing successive space-separated values into successive arguments.
	Newlines count as space.
	It returns the number of items successfully scanned. If that is less than the number of arguments, err will report why.

	func Sscan(str string, a ...interface{}) (n int, err error)
*/

func main() {
	var name string
	var age int
	n, err := fmt.Sscan("Hunter 4", &name, &age)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d: %s, %d\n", n, name, age)

}
