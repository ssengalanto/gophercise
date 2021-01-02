package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(213))
}

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
	}

	rev := 0

	for curr := x * sign; curr > 0; curr /= 10 {
		rev *= 10
		rev += curr % 10
		if rev > math.MaxInt32 {
			return 0
		}
	}

	return sign * rev
}
