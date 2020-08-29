package main

import "fmt"

func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}
	left, right := 2, num/2
	for left <= right {
		mid := left + (right-left)>>1
		squr := mid * mid
		if squr == num {
			return true
		} else if squr > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func main() {
	x := 16
	fmt.Println("Output: ", isPerfectSquare(x))
}
