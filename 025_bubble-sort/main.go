package main

import "fmt"

func main() {
	fmt.Println(bubbleSort([]int{8, 6, 4, 5, 2, 1, 7, 3}))
}

func bubbleSort(n []int) []int {
	var swp bool

	for i := len(n); i > 0; i-- {
		swp = true

		for j := 0; j < i-1; j++ {
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
				swp = false
			}
		}
		if swp {
			break
		}
	}
	return n
}
