package main

import (
	"fmt"
	"strings"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func hasCommon(w1, w2 string) bool {
	for _, ch := range w1 {
		if strings.Index(w2, string(ch)) != -1 {
			return true
		}
	}
	return false
}

func maxProduct(words []string) int {
	n := len(words)
	if n < 1 {
		return 0
	}
	result := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if !hasCommon(words[i], words[j]) {
				result = max(result, len(words[i])*len(words[j]))
			}
		}
	}

	return result
}

func main() {
	words := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	fmt.Println("Output: ", maxProduct(words))
}
