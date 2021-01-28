package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(a []int) int {
	max := -(1 << 63)
	crt := 0

	for _, v := range a {
		if crt+v < v {
			crt = v
		} else {
			crt += v
		}

		if crt > max {
			max = crt
		}
	}

	return max
}
