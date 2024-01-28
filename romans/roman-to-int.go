package main

import "fmt"

func romanToInt(s string) int {
	roman := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	ttl := 0
	v := 0

	for i := 0; i < len(s); i++ {
		v = roman[s[i]]
		if i < len(s)-1 && roman[s[i+1]] > v {
			ttl -= v
		} else {
			ttl += v
		}
	}
	return ttl
}

func main() {
	tests := []string{"III", "LVIII", "MCMXCIV"}
	for _, v := range tests {
		fmt.Println(romanToInt(v))

	}

}
