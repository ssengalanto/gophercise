package main

import (
	"fmt"
)

/*
	Sscanln is similar to Sscan, but stops scanning at a newline
	and after the final item there must be a newline or EOF.

	func Sscanln(str string, a ...interface{}) (n int, err error)
*/

func main() {
	var name string
	var alphabetCount int

	fmt.Sscanln("GFG \n 3", &name, &alphabetCount)
	fmt.Printf("name: %s, alphabet_count: %d", name, alphabetCount)

}
