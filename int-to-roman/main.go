package main

import "fmt"

func main() {
	fmt.Println(intToRoman(1994))
}

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := ""
	i := 0
	for num > 0 {
		for values[i] > num {
			i++
		}
		num -= values[i]
		result += romans[i]
	}
	return result
}
