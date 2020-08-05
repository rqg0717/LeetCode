package main

import "fmt"

// dynamic programming
func canPartition(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	sum := 0
	for _, val := range nums {
		sum += val
	}
	if sum%2 != 0 {
		return false
	}
	// sum(left)=sum(right)
	sum = sum / 2
	// dp[i][j] where 0<=i<=sum and 0<=j<=n
	dp := make([][]bool, sum+1)
	for i := 0; i <= sum; i++ {
		dp[i] = make([]bool, n)
		if i == nums[0] {
			dp[i][0] = true
		}
	}
	for i := 0; i < n; i++ {
		dp[0][i] = true
	}
	for i := 0; i <= sum; i++ {
		for j := 1; j < n; j++ {
			if i >= nums[j] {
				dp[i][j] = dp[i][j-1] || dp[i-nums[j]][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[sum][n-1]
}

func canPartition1(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	sum := 0
	for _, val := range nums {
		sum += val
	}
	if sum%2 != 0 {
		return false
	}
	// sum(left)=sum(right)
	sum = sum / 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := sum; j > 0; j-- {
			if j-nums[i-1] >= 0 {
				dp[j] = dp[j] || dp[j-nums[i-1]]
			}
		}
	}
	return dp[sum]
}

func main() {
	nums := []int{1, 5, 11, 6, 1, 1}
	fmt.Println("can the array be partitioned into equal sum subsets: ", canPartition1(nums))
}
