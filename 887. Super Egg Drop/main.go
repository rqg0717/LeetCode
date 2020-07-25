package main

import "fmt"

// broken dp[k-1][n-1]
// not broken dp[k][n-1]
// dp[k][n] = dp[k][n-1] + dp[k-1][n-1] + current floor
func superEggDrop(K int, N int) int {
	if K == 1 {
		return N
	}
	if N == 0 {
		return 0
	}
	dp := make([][]int, K+1)
	for k := 0; k <= K; k++ {
		dp[k] = make([]int, N+1)
	}
	for n := 1; n <= N; n++ {
		for k := 1; k <= K; k++ {
			dp[k][n] = dp[k][n-1] + dp[k-1][n-1] + 1
			if dp[k][n] >= N {
				return n
			}
		}
	}
	return N
}

// from https://leetcode.com/festival/
func superEggDrop1(K int, N int) int {
	var i int
	dp := make([]int, K+1)
	for i = 0; dp[K] < N; i++ {
		for j := K; j > 0; j-- {
			dp[j] += dp[j-1] + 1
		}
	}
	return i
}

// set the minimum number of moves as m
// broken: 1; not broken: 0
func superEggDrop2(K int, N int) int {
	return 0
}

func main() {
	K, N := 100, 10000
	fmt.Println("the minimum number of moves: ", superEggDrop1(K, N))
}
