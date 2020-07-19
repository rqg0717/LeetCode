package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// dynamic programming
// dp[i][k][j]
// for i <= n {
//  for k <= 2 {
//	 for j <= 1 {
//    dp[i][k][j] = max(buy, sell)
//	 }
//  }
// }
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	dp := make([][3][2]int, n)
	// init first day
	dp[0][0][0], dp[0][1][0], dp[0][2][0] = 0, 0, 0
	dp[0][0][1], dp[0][1][1], dp[0][2][1] = -prices[0], -prices[0], -prices[0]

	for i := 1; i < n; i++ {
		for k := 1; k <= 2; k++ {
			// i day, do not have any stock
			// i-1 day did not have or sold today
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			// i day, have stock
			// i-1 day had stock or bought today
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[n-1][2][0]
}

func main() {
	prices := make([]int, 0)
	prices = append(prices, 1, 2, 3, 4, 5)
	fmt.Println("max profit: ", maxProfit(prices))
}
