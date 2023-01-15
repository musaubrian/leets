package main

import "fmt"

func removeElement(nums []int, val int) int {
	j := 0
	for i, num := range nums {
		if num != val {
			if j < i {
				nums[j] = num
			}
			j++
		}
	}
	return j
}

func main() {
	var arr []int
	arr = append(arr, 0, 1, 2, 2, 3, 0, 4, 2)
	result := removeElement(arr, 2)
	fmt.Println(result)
}
