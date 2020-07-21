package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// dynamic programming
// dp[i] the maximum amount of money of i rooms
// for i <= n {
//   dp[i]= max(dp[i-1], dp[i-2]+nums[i])
// }
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp0, dp1, dp2 := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		dp0 = max(dp1, dp2+nums[i])
		dp2 = dp1
		dp1 = dp0
	}
	return dp0
}

func main() {
	nums := make([]int, 0)
	nums = append(nums, 1, 2, 3, 4, 5, 6)
	fmt.Println("max profit: ", rob(nums))
}
