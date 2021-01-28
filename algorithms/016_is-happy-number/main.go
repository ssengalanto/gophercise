package main

import "fmt"

func main() {

	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {
	c := make(map[int]bool)

	for {
		res := 0

		for n > 0 {
			m := (n % 10)
			res += (m * m)
			n /= 10
		}

		if res == 1 {
			return true
		}

		if _, ok := c[res]; ok {
			return false
		} else {
			c[res] = true
		}

		n = res
	}
}
