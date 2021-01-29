package main

import "fmt"

func main() {
	run()
}

/*
cap: returns the capacity of v, according to its type
func cap(v Type) int

Array: the number of elements in v (same as len(v)).
Pointer to array: the number of elements in *v (same as len(v)).

Slice: the maximum length the slice can reach when resliced;
if v is nil, cap(v) is zero.

Channel: the channel buffer capacity, in units of elements;
if v is nil, cap(v) is zero.
*/

func run() {
	a := make([]string, 5)
	b := make([]string, 5, 7)
	c := []int{1, 2, 3, 4, 5}

	fmt.Println("cap: ", cap(a), cap(b), cap(c))
}
