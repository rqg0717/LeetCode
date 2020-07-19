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
func maxProfit(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}
	dp0 := 0
	dp1 := -50000
	for i := 0; i < len(prices); i++ {
		temp := dp0
		// i-th day, do not have any stock
		dp0 = max(dp0, dp1+prices[i]) // i-1 day did not have/sold today
		// i-th day, have stock
		dp1 = max(dp1, temp-prices[i]-fee) // i-1 day had stock/bought today
	}
	return dp0
}

func main() {
	fee := 2
	prices := make([]int, 0)
	prices = append(prices, 1, 3, 2, 8, 4, 9)
	fmt.Println("max profit: ", maxProfit(prices, fee))
}
