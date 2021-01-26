package main

import "fmt"

func main() {
	fmt.Println(insertionSort([]int{8, 6, 4, 5, 2, 1, 7, 3}))
}

func insertionSort(n []int) []int {
	var j int

	for i := 1; i < len(n); i++ {
		current := n[i]

		for j = i - 1; j >= 0 && n[j] > current; j-- {
			n[j+1] = n[j]
		}
		n[j+1] = current
	}
	return n
}
