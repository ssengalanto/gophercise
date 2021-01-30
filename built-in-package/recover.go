package main

import (
	"fmt"
	"runtime/debug"
)

/*
	recover:  allows a program to manage behavior of a panicking goroutine.
	Executing a call to recover inside a deferred function (but not any function called by it)
	stops the panicking sequence by restoring normal execution and retrieves the error value passed to the call of panic.
	If recover is called outside the deferred function it will not stop a panicking sequence.
	In this case, or when the goroutine is not panicking, or if the argument supplied to panic was nil, recover returns nil.
	Thus the return value from recover reports whether the goroutine is panicking.

	func recover() interface{}

	Note: panic and recover can be considered similar to try-catch-finally idiom in other languages, but rarely used in Go.
*/

func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
		debug.PrintStack() // prints the call stack for debugging
	}
}

func fullName(fname *string, lname *string) {
	defer recoverFullName()
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
