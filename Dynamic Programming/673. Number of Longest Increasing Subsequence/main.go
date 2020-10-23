package main

import "fmt"

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	count := make([]int, n)
	max := 1
	dp[0] = 1
	count[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		count[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[i] == dp[j]+1 {
					count[i] += count[j]
				}
			}
			if dp[i] > max {
				max = dp[i]
			}
		}
	}
	result := 0
	for i := 0; i < len(count); i++ {
		if dp[i] == max {
			result += count[i]
		}
	}
	return result
}

func main() {
	nums := []int{2, 2, 2, 2, 2}
	fmt.Println("Output: ", findNumberOfLIS(nums))
}
