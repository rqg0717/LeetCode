package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	if n < 4 {
		return nil
	}
	results := [][]int{}
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			k := j + 1
			l := n - 1
			for k < l {
				sum := nums[i] + nums[k] + nums[j] + nums[l]
				if sum < target {
					k++
				} else if sum > target {
					l--
				} else {
					results = append(results, []int{nums[i], nums[k], nums[j], nums[l]})
					for k < l && nums[k] == nums[k+1] {
						k++
					}
					for k < l && nums[l] == nums[l-1] {
						l--
					}
					k++
					l--
				}
			}
		}
	}
	return results
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4, 4, -2}
	target := 0
	fmt.Println("Output: ", fourSum(nums, target))
}
