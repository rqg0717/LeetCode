package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	nums := []rune(s)
	n := len(nums)
	results := make([]string, numRows)
	repeat := numRows*2 - 2 // 2n-2 repeat
	for i := 0; i < n; i++ {
		mod := i % repeat
		if mod < numRows {
			results[mod] += string(nums[i])
		} else {
			results[repeat-mod] += string(nums[i])
		}
	}
	return strings.Join(results, "")
}
func main() {
	s := "PAYPALISHIRING"
	numRows := 4
	fmt.Println("output:", convert(s, numRows))
}
