package main

import "fmt"

func main() {
	run()
}

/*
	copy: copies elements from a source slice into a destination slice.
	(As a special case, it also will copy bytes from a string to a slice of bytes.)
	The source and destination may overlap.
	Copy returns the number of elements copied, which will be the minimum of len(src) and len(dst).

	func copy(dst, src []Type) int
*/

func run() {
	var a = make([]int, 3)
	ca := copy(a, []int{0, 1, 2, 3})

	var b = make([]byte, 5)
	cb := copy(b, "Hello, world!") // b == []byte("Hello")

	fmt.Println("copy: ", ca, cb)
}
