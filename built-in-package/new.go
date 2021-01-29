package main

import "fmt"

func main() {
	run()
}

/*
new: allocates memory. The first argument is a type, not a value, and the value returned is a pointer to a newly allocated zero value of that type.
func new(Type) *Type
*/

func run() {
	x := new(int)
	y := newEquivalent()

	fmt.Println(*x, *y)
}

func newEquivalent() *int {
	var x int
	return &x
}
