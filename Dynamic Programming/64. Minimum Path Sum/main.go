package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// dp[i][j] = min(dp[i-1][j],dp[i][j-1]) + grid[i][j]
func minPathSum(grid [][]int) int {
	n := len(grid)
	if n < 1 {
		return 0
	}
	dp := make([][]int, n)
	for i, row := range grid {
		dp[i] = make([]int, len(row))
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < n; i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j != 0 { // first row
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 && i != 0 { // first column
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else if i != 0 && j != 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[n-1][len(dp[n-1])-1]
}

func minPathSum1(grid [][]int) int {
	n := len(grid)
	if n < 1 {
		return 0
	}
	for i := 0; i < n; i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j != 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 && i != 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else if i != 0 && j != 0 {
				grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			}
		}
	}
	return grid[n-1][len(grid[n-1])-1]
}

func main() {
	grid := [][]int{{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1}}
	fmt.Println("Output: ", minPathSum1(grid))
}
