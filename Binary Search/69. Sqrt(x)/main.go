package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
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

func mySqrt1(x int) int {
	if x == 0 {
		return 0
	}
	guess, x0 := float64(x/2), float64(x)
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

func main() {
	x := 8
	fmt.Println("Output: ", mySqrt1(x))
}
