package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
	fmt.Println(defangIPaddrImp("1.1.1.1"))
}

// declarative
func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}

// imperative
func defangIPaddrImp(address string) string {
	var res string

	for i := 0; i < len(address); i++ {
		if address[i] == 46 {
			res += "[.]"
			continue
		}
		res += string(address[i])
	}
	return res
}
