package main

import "fmt"

// dp[i] = dp[i-1] + append(subset, nums[i])
func subsets(nums []int) [][]int {
	dp := make([][][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([][]int, 0)
	}
	dp[0] = [][]int{{}}
	for i := 1; i <= len(nums); i++ {
		dp[i] = append(dp[i], dp[i-1]...)
		subset1 := make([][]int, len(dp[i-1]))
		copy(subset1, dp[i-1])
		for _, val := range subset1 {
			subset2 := make([]int, len(val))
			copy(subset2, val)
			subset2 = append(subset2, nums[i-1])
			dp[i] = append(dp[i], subset2)
		}
	}
	return dp[len(nums)]
}

// backtracking algorithm
func backtracking(i int, nums, subset1 []int, results *[][]int) {
	subset2 := make([]int, len(subset1))
	copy(subset2, subset1)
	*results = append(*results, subset2)
	for j := i; j < len(nums); j++ {
		subset1 = append(subset1, nums[j])
		backtracking(j+1, nums, subset1, results)
		subset1 = subset1[:len(subset1)-1]
	}
}

func subsets1(nums []int) [][]int {
	var results [][]int
	subset := []int{}
	backtracking(0, nums, subset, &results)
	return results
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println("Output: ", subsets1(nums))
}
