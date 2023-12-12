package main

import (
	"fmt"
	"slices"
	"sort"
)

//	func bubbleSort(nums []int) []int {
//		swapped := true
//		for swapped {
//			swapped = false
//			for i := 1; i < len(nums); i++ {
//				if nums[i-1] > nums[i] {
//					nums[i], nums[i-1] = nums[i-1], nums[i]
//					swapped = true
//				}
//			}
//		}
//		return nums
//	}

func maxProduct(nums []int) int {
	slices.Sort(nums)
	return (nums[len(nums)-1] - 1) * (nums[len(nums)-2] - 1)
}

func main() {
	fmt.Println(maxProduct([]int{3, 4, 5, 2}))
	fmt.Println(maxProduct([]int{1, 5, 4, 5}))
	fmt.Println(maxProduct([]int{3, 7}))
}
