package main

import (
	"fmt"
	"math"
)

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

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	guess, x0 := float64(x), float64(x)
	for {
		// 2xix0-(x0^2+guess) = 0
		// xi = 1/2(x0+guess/x0)
		xi := 0.5 * (x0 + guess/x0)
		if math.Abs(x0-xi) < 1e-7 {
			break
		}
		x0 = xi
	}
	return int(x0)
}

func bulbSwitch1(n int) int {
	return mySqrt(n)
}

func main() {
	n := 16
	fmt.Println("Output: ", bulbSwitch1(n))
}
