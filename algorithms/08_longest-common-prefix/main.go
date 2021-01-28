package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{
		"flower",
		"flow",
		"flight",
	}))
}

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for _, v := range strs {
		if v == "" || prefix == "" {
			return ""
		}

		if len(prefix) < len(v) {
			v = v[:len(prefix)]
		} else {
			prefix = prefix[:len(v)]
		}

		for i := 0; i < len(v); i++ {
			if prefix[i] != v[i] {
				prefix = v[0:i]
				break
			}
		}
	}
	return prefix
}
