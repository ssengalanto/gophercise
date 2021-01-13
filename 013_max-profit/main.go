package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max, min := 0, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[min] {
			min = i
		} else if cur := prices[i] - prices[min]; cur > max {
			max = cur
		}
	}
	return max
}
