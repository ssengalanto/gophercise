package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

var romanMap = map[string]int{
	"I":  1,
	"V":  5,
	"X":  10,
	"L":  50,
	"C":  100,
	"D":  500,
	"M":  1000,
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}

	var res int
	idx := 0
	for idx < len(s) {
		if idx+1 < len(s) && (s[idx] == 'I' || s[idx] == 'X' || s[idx] == 'C') {
			sub := s[idx : idx+2]
			if v, ok := romanMap[sub]; ok {
				res += v
				idx += 2
				continue
			}
		}

		sub := s[idx : idx+1]
		res += romanMap[sub]
		idx++
	}
	return res
}
