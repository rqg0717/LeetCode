package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSubArray(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	max := nums[0]
	if n == 1 {
		return max
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i]+dp[i-1] > nums[i] {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("Output: ", maxSubArray(nums))
}
