package main

import (
	"fmt"
)

func circularPermutation(n int, start int) []int {
	results := []int{start}
	count := 1
	for i := 0; i < n; i++ {
		for j := count - 1; j >= 0; j-- {
			results = append(results, results[j]^count)
		}
		count <<= 1
	}
	return results
}

func main() {
	n := 2
	start := 3
	fmt.Println("Output: ", circularPermutation(n, start))
}
