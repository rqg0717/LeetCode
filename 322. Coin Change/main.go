package main

import "fmt"

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

/*
	e.g. coins = [1, 2, 5], amount = 11
	then dp(11) = min(dp(11 - 1), dp(11 - 2), dp(11 - 5)) + 1
	---> dp(amount) = min(dp(amount - coins[len(coins)])) + 1
*/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				dp[i] = min(dp[i-coins[j]]+1, dp[i])
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

func main() {
	coins := []int{1, 5, 25, 100}
	amount := 113
	fmt.Println("output: ", coinChange(coins, amount))
}
