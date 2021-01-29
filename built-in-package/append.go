package main

import "fmt"

func main() {
	run()
}

func run() {
	a := []int{4, 5}
	b := []int{1, 2, 3}

	// append: appends elements to the end of a slice, it returns the updated slice so its result is usually stored in a variable
	// func append(slice []Type, elems ...Type) []Type
	ab := append(a, b...)

	fmt.Println(a, b, ab)
}
