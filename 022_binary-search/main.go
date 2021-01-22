package main

import "fmt"

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

func search(nums []int, target int) int {
	for i, j := 0, len(nums)-1; i <= j; {
		m := i + (j-i)/2
		if nums[m] < target {
			i = m + 1
		} else if nums[m] > target {
			j = m - 1
		} else {
			return m
		}
	}
	return -1
}
