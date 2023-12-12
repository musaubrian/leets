package main

import (
	"fmt"
	"strings"
)

// Return the first index when a substring is found in a string
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
	fmt.Println(strStr("leetcode", "leeto"))
	fmt.Println(strStr("aaa", "aaa"))
}
