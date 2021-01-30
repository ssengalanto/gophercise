package main

import "fmt"

func main() {
	run()
}

/*
	append: appends elements to the end of a slice. If it has sufficient capacity,
	the destination is resliced to accommodate the new elements.
	If it does not, a new underlying array will be allocated.
	Append returns the updated slice.
	It is therefore necessary to store the result of append, often in the variable holding the slice itself.

	func append(slice []Type, elems ...Type) []Type
*/

func run() {
	a := []int{4, 5}
	b := []int{1, 2, 3}

	ab := append(a, b...)

	fmt.Println(a, b, ab)
}
