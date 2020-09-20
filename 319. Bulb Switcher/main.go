package main

import "fmt"

func bulbSwitch(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		count := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count%2 != 0 { // on when odd
			result++
		}
	}
	return result
}

func sqrt(x int) int {
	left, right := 0, x
	ans := -1
	for left <= right {
		mid := left + (right-left)>>1
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

func bulbSwitch1(n int) int {
	return sqrt(n)
}

func main() {
	n := 16
	fmt.Println("Output: ", bulbSwitch1(n))
}
