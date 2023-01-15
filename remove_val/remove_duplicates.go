package main

func removeDuplicates(nums []int) int {
	j := 0
	for i, num := range nums {
		if i == 0 || num != nums[i-1] {
			nums[j] = num
			j += 1
		}
	}
	return j
}
