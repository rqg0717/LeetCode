package main

import "fmt"

// broken dp[k-1][n-1]
// not broken dp[k][n-1]
// dp[k][n] = dp[k][n-1] + dp[k-1][n-1] + current floor
func superEggDrop(K int, N int) int {
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

// set the minimum number of moves as m
// broken: 1; not broken: 0
func superEggDrop1(K int, N int) int {
	return 0
}

func main() {
	K, N := 1, 2
	fmt.Println("the minimum number of moves: ", superEggDrop1(K, N))
}
