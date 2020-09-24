package main

import "fmt"

func minSteps(n int) int {
	result := 0
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			result += i
			n /= i
		}
	}
	return result
}

func minSteps1(n int) int {
	if n < 2 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	dp[2] = 2 // copy and paste
	for i := 3; i < n+1; i++ {
		steps := i
		for j := i / 2; j >= 2; j-- {
			if i%j == 0 {
				steps = dp[j] + i/j
				break
			}
		}
		dp[i] = steps
	}
	return dp[n]
}

func main() {
	fmt.Println("Output: ", minSteps1(30))
}
