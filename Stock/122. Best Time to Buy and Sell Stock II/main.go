package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// dynamic programming
// dp[i][j]
// for i <= n {
//  for j <= 1 {
//   dp[i][j] = max(buy, sell)
//  }
// }
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp0 := 0
	dp1 := -prices[0]
	for i := 0; i < len(prices); i++ {
		temp := dp0
		// today, do not have any stock
		dp0 = max(dp0, dp1+prices[i]) // yesterday did not have/sold today
		// today, have stock
		dp1 = max(dp1, temp-prices[i]) // yesterday had stock/bought today
	}
	return dp0
}

func main() {
	prices := make([]int, 0)
	prices = append(prices, 7, 1, 5, 3, 6, 4)
	fmt.Println("max profit: ", maxProfit(prices))
}
