package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	dp := make([]int, n)
	count := 1
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > count {
			count = dp[i]
		}
	}
	result := 0
	for i := 0; i < n; i++ {
		if dp[i] == count {
			result++
		}
	}
	return result
}

func main() {
	nums := []int{1, 3, 5, 4, 7}
	fmt.Println("Output: ", findNumberOfLIS(nums))
}
