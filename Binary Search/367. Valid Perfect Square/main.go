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

// Newton–Raphson method
func isPerfectSquare1(num int) bool {
	if num < 2 {
		return true
	}
	guess := num >> 1
	for guess*guess > num {
		guess = (guess + num/guess) >> 1
	}
	return (guess*guess == num)
}

func main() {
	x := 16
	fmt.Println("Output: ", isPerfectSquare1(x))
}
