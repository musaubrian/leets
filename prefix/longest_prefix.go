package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	if len(strs) <= 1 {
		return prefix
	}

	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if len(strs[i]) <= j || strs[i][j] != prefix[j] {
				prefix = prefix[:j]
				break
			}
		}
	}

	return prefix
}

func main() {
	tests := map[int][]string{
		0: {"flower", "flow", "floght"},
		1: {"dog", "racecar", "car"},
	}

	for _, v := range tests {
		fmt.Println(longestCommonPrefix(v))
	}
}
