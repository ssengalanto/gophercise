package main

import "fmt"

func main() {
	run()
}

/*
make: allocates and initializes an object of type slice, map, or chan (only)
func make(t Type, size ...IntegerType) Type

Slice: The size specifies the length. The capacity of the slice is
equal to its length. A second integer argument may be provided to
specify a different capacity; it must be no smaller than the
length. For example, make([]int, 0, 10) allocates an underlying array
of size 10 and returns a slice of length 0 and capacity 10 that is
backed by this underlying array.

Map: An empty map is allocated with enough space to hold the
specified number of elements. The size may be omitted, in which case
a small starting size is allocated.

Channel: The channel's buffer is initialized with the specified
buffer capacity. If zero, or the size is omitted, the channel is
unbuffered.*/

func run() {

	a := make([]int, 5)
	b := make(map[string]int)
	c := make(chan string)

	fmt.Println(a, b, c)
}
