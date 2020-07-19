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
//  for k <= K {
//	 for j <= 1 {
//    dp[i][k][j] = max(buy, sell)
//	 }
//  }
// }
func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 || k < 1 {
		return 0
	}

	// the maximum transactions
	if k >= n/2 {
		profit := 0
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				profit += prices[i] - prices[i-1]
			}
		}
		return profit
	}

	dp := make([][2]int, k+1)
	// init first day
	for i := 0; i <= k; i++ {
		dp[i][0] = 0
		dp[i][1] = -prices[0]
	}

	for i := 0; i < n; i++ {
		for kk := 1; kk <= k; kk++ {
			// i day, do not have any stock
			// i-1 day did not have or sold today
			dp[kk][0] = max(dp[kk][0], dp[kk][1]+prices[i])
			// i day, have stock
			// i-1 day had stock or bought today
			dp[kk][1] = max(dp[kk][1], dp[kk-1][0]-prices[i])
		}
	}
	return dp[k][0]
}

func main() {
	k := 2
	prices := make([]int, 0)
	prices = append(prices, 3, 3, 5, 0, 0, 3, 1, 4)
	fmt.Println("max profit: ", maxProfit(k, prices))
}
