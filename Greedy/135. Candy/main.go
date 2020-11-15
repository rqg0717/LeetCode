package main

import (
	"fmt"
)

func candy(ratings []int) int {
	n := len(ratings)
	if n < 2 {
		return n
	}
	nums := make([]int, n)

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			nums[i] = nums[i-1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && nums[i] <= nums[i+1] {
			nums[i] = nums[i+1] + 1
		}
	}

	sum := n
	for _, v := range nums {
		sum += v
	}

	return sum
}

func main() {
	ratings := []int{1, 0, 2}
	fmt.Println("Sum: ", candy(ratings))
}
