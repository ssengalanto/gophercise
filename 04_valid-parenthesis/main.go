package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("([])"))
}

func isValid(s string) bool {
	if s == "" {
		return false
	}

	stack := []rune{}
	bracketsMaps := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, v := range s {
		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		}
		if bracketsMaps[v] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return len(stack) == 0
}
