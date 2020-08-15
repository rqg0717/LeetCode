package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return n
	}
	m := map[byte]int{}
	count, left, right := 0, 0, 0
	for right < n {
		m[s[right]]++
		for m[s[right]] > 1 {
			m[s[left]]--
			left++

		}
		right++
		if (right - left) > count {
			count = right - left
		}

	}
	return count
}

func main() {
	s := "abcabcbb"
	fmt.Println("Output: ", lengthOfLongestSubstring(s))
}
