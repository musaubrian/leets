package main

import "math"

func binary(nums []int, target int) int {
	var (
		low  = 0
		high = len(nums)
	)

	for low < high {
		midpoint := int(math.Floor(float64(low) + (float64(high)-float64(low))/2))
		val := nums[midpoint]
		if target == val {
			return midpoint
		} else if val < target {
			low = midpoint + 1
		} else {
			high = midpoint
		}
	}
	return -1
}
