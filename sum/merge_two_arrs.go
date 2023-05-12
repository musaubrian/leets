package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	// makes runtime slower*
	fmt.Println(nums1)
}

func main() {
	var nums []int
	var nums2 []int
	nums = append(nums, 1, 2, 3, 0, 0, 0)
	nums2 = append(nums2, 2, 5, 6)
	merge(nums, 3, nums2, 3)
}
