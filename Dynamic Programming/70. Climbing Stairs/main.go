package main

import (
	"fmt"
)

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbStairs1(n int) int {
	i, j, k := 0, 0, 1
	for l := 1; l <= n; l++ {
		i = j
		j = k
		k = i + j
	}
	return k
}

func main() {
	fmt.Println("Output: ", climbStairs1(3))
}
