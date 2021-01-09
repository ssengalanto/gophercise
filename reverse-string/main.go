package main

import "fmt"

func main() {
	fmt.Println(reverseString([]string{"h", "e", "l", "l", "o"}))
}

func reverseString(s []string) []string {
	i, j := 0, len(s)-1

	for i <= j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}

	return s
}
