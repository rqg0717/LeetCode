package main

import (
	"fmt"
	"sort"
)

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	// between [0, max(nums) - min(nums)]
	left, right := 0, nums[len(nums)-1]-nums[0]
	for left < right {
		mid := (left + right) >> 1
		i := 0
		counter := 0
		for j := range nums {
			for nums[j]-nums[i] > mid {
				i++
			}
			counter = counter + j - i
		}
		if counter >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	nums := []int{0, 1, 2, 4, 6, 5}
	k := 5
	fmt.Println("Output: ", smallestDistancePair(nums, k))
}
