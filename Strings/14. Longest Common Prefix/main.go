package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for _, str := range strs {
		for strings.Index(str, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

func longestCommonPrefix1(strs []string) string {
	if len(strs) < 1 {
		return ""
	}

	isCommonPrefix := func(length int) bool {
		str0, count := strs[0][:length], len(strs)
		for i := 1; i < count; i++ {
			if strs[i][:length] != str0 {
				return false
			}
		}
		return true
	}
	minLength := len(strs[0])
	for _, str := range strs {
		if len(str) < minLength {
			minLength = len(str)
		}
	}
	low, high := 0, minLength
	for low < high {
		mid := (high-low+1)/2 + low
		if isCommonPrefix(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return strs[0][:low]
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println("Output: ", longestCommonPrefix1(strs))
}
