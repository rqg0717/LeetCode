package main

import (
	"fmt"
	"sort"
)

func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return []int{nums[1], nums[0]}
		}
		return nums
	}
	var low, high []int
	for _, n := range nums[1:] {
		if n < nums[0] {
			low = append(low, n)
		} else {
			high = append(high, n)
		}
	}
	var rerults = append(sortArray(low), nums[0])
	rerults = append(rerults, sortArray(high)...)
	return rerults
}

// double pointers
func threeSum(nums []int) [][]int {
	n := len(nums)
	//sortedNums := sortArray(nums)
	sort.Ints(nums)
	results := [][]int{}
	for i := 0; i < n; i++ {
		target := 0 - nums[i]
		left := i + 1
		right := n - 1
		if nums[i] > 0 {
			break
		}
		if i == 0 || nums[i] != nums[i-1] {
			for left < right {
				if nums[left]+nums[right] == target {
					results = append(results, []int{nums[i], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if nums[left]+nums[right] < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return results
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4, 4, 2}
	fmt.Println("Output: ", threeSum(nums))
}
