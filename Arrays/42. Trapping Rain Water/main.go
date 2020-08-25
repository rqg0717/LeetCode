package main

import (
	"fmt"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func trap(height []int) int {
	n := len(height)
	result := 0
	for i := 1; i < n-1; i++ {
		leftMax := 0
		for j := 0; j <= i; j++ {
			leftMax = max(height[j], leftMax)
		}
		rightMax := 0
		for j := i; j < n; j++ {
			rightMax = max(height[j], rightMax)
		}
		result = result + min(leftMax, rightMax) - height[i]
	}
	return result
}

func trap1(height []int) int {
	right := len(height) - 1
	left, leftMax, rightMax, result := 0, 0, 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				result += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				result += rightMax - height[right]
			}
			right--
		}
	}
	return result
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println("Output: ", trap1(height))
}
