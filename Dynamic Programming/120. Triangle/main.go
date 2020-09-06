package main

import (
	"fmt"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func dfs(i, j, min, sum int, triangle [][]int) int {
	if i == len(triangle) {
		if min > sum {
			min = sum
		}
		return min
	}
	a := dfs(i+1, j, min, sum+triangle[i][j], triangle)
	b := dfs(i+1, j+1, min, sum+triangle[i][j], triangle)
	if a < b {
		return a
	}
	return b
}

func minimumTotal(triangle [][]int) int {
	i, j, min, sum := 0, 0, 0, 0
	for n := 0; n < len(triangle); n++ {
		min = min + triangle[n][0]
	}
	return dfs(i, j, min, sum, triangle)
}

// dp
func minimumTotal1(triangle [][]int) int {
	n := len(triangle)
	if n == 1 {
		return triangle[0][0]
	}
	dp := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			dp = triangle[i-1][j] + triangle[i][j]
			if j-1 >= 0 {
				dp = min(dp, triangle[i-1][j-1]+triangle[i][j])
			}
			triangle[i][j] = dp
		}
		triangle[i][i] += triangle[i-1][i-1]
	}
	dp = triangle[n-1][0]
	for _, val := range triangle[n-1] {
		if val < dp {
			dp = val
		}
	}
	return dp
}

func minimumTotal2(triangle [][]int) int {
	n := len(triangle)
	if n == 1 {
		return triangle[0][0]
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if dp[i] == nil {
				dp[i] = make([]int, len(triangle[i]))
			}
			dp[i][j] = triangle[i][j]
		}
	}
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
		}
	}
	return dp[0][0]
}

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println("Output: ", minimumTotal2(triangle))
}
