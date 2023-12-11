package main

import (
	"fmt"
)

func isValid(s string) bool {
	opts := map[byte]byte{'(': ')', '{': '}', '[': ']'}
	pairs := []byte{}

	for _, r := range []byte(s) {
		if close, ok := opts[r]; ok {
			pairs = append(pairs, close)
		} else {
			if len(pairs) == 0 || pairs[len(pairs)-1] != r {
				return false
			}
			pairs = pairs[:len(pairs)-1]
		}
	}
	return len(pairs) == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("()"))
	fmt.Println(isValid("(]"))

}
