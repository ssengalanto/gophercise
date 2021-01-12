package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	max := 0
	i, j := 0, len(height)-1
	var vol int
	for i < j {
		if height[i] > height[j] {
			vol = (j - i) * height[j]
			j--
		} else {
			vol = (j - i) * height[i]
			i++
		}
		if max < vol {
			max = vol
		}
	}
	return max
}
