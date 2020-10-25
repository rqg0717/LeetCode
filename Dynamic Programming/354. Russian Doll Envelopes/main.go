package main

import (
	"fmt"
	"sort"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n <= 1 {
		return n
	}

	// sort array
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	dp := make([]int, n)
	dp[0] = 1
	result := 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[j][1] < envelopes[i][1] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		result = max(result, dp[i])
	}
	return result
}

func main() {
	envelopes := [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	fmt.Println("output: ", maxEnvelopes(envelopes))
}
