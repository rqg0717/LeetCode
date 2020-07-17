package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	buy := -prices[0]
	profit := 0

	for i := 1; i < n; i++ {
		timetoBuy := max(buy, -prices[i])
		timetoSell := max(buy+prices[i], profit)
		buy, profit = timetoBuy, timetoSell
	}
	return profit
}

func main() {
	prices := make([]int, 0)
	prices = append(prices, 7, 1, 5, 3, 6, 4)
	fmt.Println("max profit: ", maxProfit(prices))
}
