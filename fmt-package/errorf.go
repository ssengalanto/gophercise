package main

import (
	"fmt"
)

/*
	Errorf formats according to a format specifier and returns the string as a value that satisfies error.

	If the format specifier includes a %w verb with an error operand, the returned error will implement an Unwrap method returning the operand.
	It is invalid to include more than one %w verb or to supply it with an operand that does not implement the error interface.
	The %w verb is otherwise a synonym for %v.
*/

func main() {
	const name, id = "Hunter", 4
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	err2 := fmt.Errorf("user %w (id %w) not found", name, id)
	fmt.Println(err.Error())
	fmt.Println(err2.Error())
}
