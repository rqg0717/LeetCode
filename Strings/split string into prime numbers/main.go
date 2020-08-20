package main

import (
	"fmt"
	"math"
	"strconv"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func countPrimes(s string) int {
	n := int(math.Pow10(len(s)))
	// Sieve of Eratosthenes
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	l := len(s) + 1
	dp := make([]int, l)
	for i := 1; i < l; i++ {
		count := 0
		for j := 1; j < i; j++ {
			index, _ := strconv.Atoi(s[j:i])
			if index > n {
				break
			}
			if index > 1 {
				if isPrime[index] {
					count = count + dp[j]
				}
			}
		}
		dp[i] = count
		index, _ := strconv.Atoi(s[:i])
		if index > n {
			continue
		}
		if index > 1 {
			if isPrime[index] {
				dp[i]++
			}
		}
	}
	return dp[l-1]
}

func main() {
	s := "3175"
	fmt.Println("output:", countPrimes(s))
}
