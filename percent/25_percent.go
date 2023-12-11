package main

import (
	"fmt"
)

func findSpecialInteger(arr []int) int {
	val := 0
	dict := make(map[int]int)
	for _, r := range arr {
		dict[r]++
	}
	for k, v := range dict {
		if v > len(arr)/4 {
			val = k
		}
	}
	return val
}

func main() {
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10}))
}
