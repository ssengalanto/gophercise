package main

import "fmt"

func main() {
	fmt.Println(fib(10))
}

func fib(n int) int {
	fibs := map[int]int{1: 1, 2: 1}

	for i := 3; i <= n; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}

	return fibs[n]
}
