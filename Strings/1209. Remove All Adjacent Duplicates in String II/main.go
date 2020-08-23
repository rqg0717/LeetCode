package main

import (
	"fmt"
)

func removeDuplicates(s string, k int) string {
	n := len(s)
	if n < k {
		return s
	}
	letters := []rune(s)
	count := make([]int, n)
	for i := 0; i < n; i++ {
		if i == len(letters) {
			return string(letters)
		}
		if i == 0 || letters[i] != letters[i-1] {
			count[i] = 1
			continue
		}

		count[i] = count[i-1] + 1
		if count[i] == k {
			letters = append(letters[:i-k+1], letters[i+1:]...)
			i = i - k
		}
	}
	return string(letters)
}

func main() {
	s := "pbbcggttciiippooaais"
	k := 2
	fmt.Println("Output: ", removeDuplicates(s, k))
}
