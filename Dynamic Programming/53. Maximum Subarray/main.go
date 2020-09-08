package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSubArray(nums []int) int {

}

func main() {
	nums := []int{10, 9, 4, 5, 2, 7, 1, 0, 6}
	fmt.Println("Output: ", maxSubArray(nums))
}
