package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 1
	ret := 0
	for i := 1; i < n; i++ {
		value := 0
		for j := 0; j < i; j++ {
			if nums[j] > nums[i] {
				value = max(value, dp[j])
			}
		}
		dp[i] = value + 1
		ret = max(ret, dp[i])
	}
	return ret
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println("Output: ", lengthOfLIS(nums))
}
