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

func main() {
	fmt.Println("Output: ", minSteps(3))
}
