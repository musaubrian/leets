package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
	res := math.Sqrt(float64(x))
	return int(math.Floor(res))
}
func main() {
	test := []int{4, 8, 0}
	for _, v := range test {
		fmt.Println(mySqrt(v))
	}
}
