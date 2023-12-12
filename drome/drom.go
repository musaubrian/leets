package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y := x
	// reverse a number
	res := 0
	for y != 0 {
		res = res*10 + y%10
		y /= 10
	}

	return res == x
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(1000021))
}
