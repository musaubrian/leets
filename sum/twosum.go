package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}

func main() {
	var arr []int
	arr = append(arr, 3, 5, 8, 11, 15, 18, 20, 25, 30, 35)
	result := twoSum(arr, 38)
	fmt.Println(result)
}
