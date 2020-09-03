package main

import (
	"fmt"
	"math"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func splitArray(nums []int, m int) int {
	left, right := 0, 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if left < nums[i] {
			left = nums[i]
		}
		right += nums[i]
	}
	for left < right {
		mid := left + (right-left)>>1
		if checkSum(nums, mid, m, n) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func checkSum(nums []int, mid, m, n int) bool {
	sum, counter := 0, 1
	for i := 0; i < n; i++ {
		if sum+nums[i] > mid {
			counter++
			sum = nums[i]
		} else {
			sum += nums[i]
		}
	}
	return counter <= m
}

// minimize the largest sum among these m subarrays
// dp[i][j]=min{max(dp[k][j−1],sum(k+1,i))}
func splitArray1(nums []int, m int) int {
	n := len(nums)
	sum := make([]int, n+1)
	// sum(i, j)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	dp := make([][]int, n+1)
	// init
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	// k != 0
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= min(i, m); j++ {
			for k := 0; k < i; k++ {
				dp[i][j] = min(dp[i][j], max(dp[k][j-1], sum[i]-sum[k]))
			}
		}
	}
	return dp[n][m]
}

func main() {
	nums := []int{0, 1, 2, 4, 6, 5, 1}
	m := 2
	fmt.Println("Output: ", splitArray1(nums, m))
}
