package main

import (
	"fmt"
)

func grayCode(n int) []int {
	results := []int{0}
	count := 1
	for i := 0; i < n; i++ {
		for j := count - 1; j >= 0; j-- {
			results = append(results, results[j]+count)
		}
		count <<= 1
	}
	return results
}

func main() {
	n := 7
	fmt.Println("Output: ", grayCode(n))
}
