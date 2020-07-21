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
// three cases: 1. do not rob the first room and the last room
//				2. rob the first room then do not rob the last room
//				3. rob the last room then do not rob the first room
func rob(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	profit := max(robRooms(nums, 0, n-1),
		robRooms(nums, 1, n))
	profit = max(profit, robRooms(nums, 1, n-1))
	return profit
}

func robRooms(nums []int, start int, end int) int {
	dp0, dp1, dp2 := 0, 0, 0
	for i := start; i < end; i++ {
		dp0 = max(dp1, dp2+nums[i])
		dp2 = dp1
		dp1 = dp0
	}
	return dp0
}

func main() {
	nums := make([]int, 0)
	nums = append(nums, 2, 3, 2)
	fmt.Println("max profit: ", rob(nums))
}
