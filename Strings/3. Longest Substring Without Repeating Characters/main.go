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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLongestSubstring1(s string) int {
	n := len(s)
	count := 0
	var ASCII [256]int // total number of chars in ASCII table is 256
	for i, j := 0, 0; i < n; i++ {
		j = max(ASCII[s[i]], j)
		count = max(count, i-j+1)
		ASCII[s[i]] = i + 1
	}
	return count
}

func main() {
	s := "a"
	fmt.Println("Output: ", lengthOfLongestSubstring1(s))
}
