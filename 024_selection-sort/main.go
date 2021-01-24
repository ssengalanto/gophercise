package main

import "fmt"

func main() {
	fmt.Println(selectionSort([]int{8, 6, 4, 5, 2, 1, 7, 3}))
}

func selectionSort(n []int) []int {
	for i := 0; i < len(n); i++ {
		min := i
		for j := i + 1; j < len(n); j++ {
			if n[j] < n[min] {
				min = j
			}
		}

		if i != min {
			n[i], n[min] = n[min], n[i]
		}
	}

	return n
}
