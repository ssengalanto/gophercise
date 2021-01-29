package main

import "fmt"

func main() {
	run()
}

/*
	len: returns the length of v, according to its type:
	func len(v Type) int

	Array: the number of elements in v.

	Pointer to array: the number of elements in *v (even if v is nil).

	Slice, or map: the number of elements in v; if v is nil, len(v) is zero.

	String: the number of bytes in v.

	Channel: the number of elements queued (unread) in the channel buffer;
         if v is nil, len(v) is zero.
*/

func run() {
	a := make([]string, 5)
	b := make([]string, 5, 7)
	c := []int{1, 2, 3, 4, 5}

	fmt.Println("len: ", len(a), len(b), len(c))
}
