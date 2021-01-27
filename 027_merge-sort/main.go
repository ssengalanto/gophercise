package main

import "fmt"

func main() {
	fmt.Println(mergeSort([]int{8, 6, 4, 5, 2, 1, 7, 3}))
}

func mergeSort(n []int) []int {
	if len(n) <= 1 {
		return n
	}
	mid := len(n) / 2
	left := mergeSort(n[:mid])
	right := mergeSort(n[mid:])
	return merge(left, right)
}

func merge(x, y []int) []int {
	i, j := 0, 0
	var res []int

	for i < len(x) && j < len(y) {
		if x[i] < y[j] {
			res = append(res, x[i])
			i++
		} else {
			res = append(res, y[j])
			j++
		}
	}

	for i < len(x) {
		res = append(res, x[i])
		i++
	}

	for j < len(y) {
		res = append(res, y[j])
		j++
	}

	return res
}
