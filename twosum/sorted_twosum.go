package main

import (
	"fmt"
)

func sortedTwoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum < target {
			i++
		} else {
			j--
		}
	}

	return nil
}
func main() {
	var arr []int
	arr = append(arr, 3, 5, 8, 11, 15, 18, 20, 25, 30, 35)
	result := sortedTwoSum(arr, 38)
	fmt.Println(result)
}
