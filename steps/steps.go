package main

import "fmt"

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	prev, curr := 1, 1

	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}

	return curr
	// recursive approach was way slower, like a lot
	// return climbStairs(n-1) + climbStairs(n-2)
}

func main() {
	tests := []int{2, 3, 5, 6, 45}
	for i := 0; i < len(tests); i++ {
		fmt.Println(climbStairs(tests[i]))
	}

}
