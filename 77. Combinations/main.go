package main

import "fmt"

func combine(n int, k int) [][]int {
	results := [][]int{}
	backtracking(1, k, n, []int{}, &results)
	return results
}

func backtracking(start, k, n int, subset []int, results *[][]int) {
	if len(subset) == k {
		combination := make([]int, k)
		copy(combination, subset)
		*results = append(*results, combination)
		return // results
	}
	for i := start; i <= n; i++ {
		subset = append(subset, i)
		backtracking(i+1, k, n, subset, results)
		subset = subset[:len(subset)-1]
	}
}

func main() {
	n := 4
	k := 2
	fmt.Println("Output: ", combine(n, k))
}
