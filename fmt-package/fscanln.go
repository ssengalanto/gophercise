package main

import (
	"fmt"
	"io"
	"strings"
)

/*
	Fscanln is similar to Fscan, but stops scanning at a newline and after the final item there must be a newline or EOF.

	func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
*/

func main() {
	s := `dmr 1771 1.61803398875
	ken 271828 3.14159`
	r := strings.NewReader(s)
	var a string
	var b int
	var c float64
	for {
		n, err := fmt.Fscanln(r, &a, &b, &c)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d: %s, %d, %f\n", n, a, b, c)
	}
}
