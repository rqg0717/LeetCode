package main

import (
	"fmt"
)

func numSteps(s string) int {
	n := len(s)
	i, zeros := 0, 0
	for j := 1; j < n; j++ {
		if s[j] == '1' {
			zeros += j - i - 1
			i = j
		}
	}
	if i == 0 {
		return n - 1
	}
	return zeros + 1 + n
}

func main() {
	s := "1101"
	fmt.Println("Output: ", numSteps(s))
}
