package main

import "fmt"

/*
Printf formats according to a format specifier and writes to standard output.
It returns the number of bytes written and any write error encountered.

func Printf(format string, a ...interface{}) (n int, err error)
*/

func main() {
	const name, age = "Hunter", 4
	p, err := fmt.Printf("%s is %d years old.\n", name, age)
	fmt.Println(p, err)
}
