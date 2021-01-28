package main

import "fmt"

func main() {
	fmt.Println(countPrimes(10))
}

func countPrimes(n int) int {

	if n == 0 || n == 1 {
		return 0
	}

	var count int

	notPrime := make([]bool, n)
	notPrime[0], notPrime[1] = true, true

	for i := 2; i < n; i++ {
		if notPrime[i] {
			continue
		}
		count++
		for j := i * 2; j < n; j += i {
			notPrime[j] = true
		}
	}

	return count
}
