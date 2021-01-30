package main

import (
	"fmt"
)

/*
	panic:  stops normal execution of the current goroutine.
	When a function F calls panic, normal execution of F stops immediately.
	Any functions whose execution was deferred by F are run in the usual way, and then F returns to its caller.
	To the caller G, the invocation of F then behaves like a call to panic, terminating G's execution and running any deferred functions.
	This continues until all functions in the executing goroutine have stopped, in reverse order.
	At that point, the program is terminated with a non-zero exit code.
	This termination sequence is called panicking and can be controlled by the built-in function recover.

	func panic(v interface{})

	Note: panic and recover can be considered similar to try-catch-finally idiom in other languages, but rarely used in Go.
*/

func fullName(fname *string, lname *string) {
	defer fmt.Println("deferred call before panic in fullName call")
	if fname == nil {
		panic("runtime error: fname cannot be nil")
	}
	if lname == nil {
		panic("runtime error: lname cannot be nil")
	}
	defer fmt.Println("deferred call after panic in fullName call")

	fmt.Printf("%s %s\n", *fname, *lname)
	fmt.Println("returned normally from fullName")
}

func main() {
	defer fmt.Println("deferred call in main")
	fname := "Ssen"
	fullName(&fname, nil)
	fmt.Println("returned normally from main")
}
