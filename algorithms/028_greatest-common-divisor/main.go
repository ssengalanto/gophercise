package main

import "fmt"

func main() {
	fmt.Println(getGCD(5, 10))
}

func getGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
