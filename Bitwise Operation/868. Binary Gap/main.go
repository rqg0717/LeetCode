package main

import (
	"fmt"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func binaryGap(n int) int {
	result := 0
	last := -1
	for i := 0; i < 32; i++ {
		if (n>>i)&1 > 0 {
			if last >= 0 {
				result = max(result, i-last)
			}
			last = i
		}
	}
	return result
}

func main() {
	n := 22
	fmt.Println("Output: ", binaryGap(n))
}
