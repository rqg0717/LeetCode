package main

import (
	"fmt"
	"sort"
)

// quick sort
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
	low, high := []int{}, []int{}
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

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	sortedNums := sortArray(nums)
	results := [][]int{}
	for i := 0; i < n; i++ {
		target := 0 - sortedNums[i]
		left := i + 1
		right := n - 1
		if sortedNums[i] > 0 {
			break
		}
		if i == 0 || sortedNums[i] != sortedNums[i-1] {
			for left < right {
				if sortedNums[left]+sortedNums[right] == target {
					results = append(results, []int{sortedNums[i], sortedNums[left], sortedNums[right]})
					for left < right && sortedNums[left] == sortedNums[left+1] {
						left++
					}
					for left < right && sortedNums[right] == sortedNums[right-1] {
						right--
					}
					left++
					right--
				} else if sortedNums[left]+sortedNums[right] < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return results
}

func fourSum1(nums []int, target int) [][]int {
	n := len(nums)
	if n < 4 {
		return nil
	}
	results := [][]int{}
	sort.Ints(nums)
	//fmt.Println(nums)
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
	fmt.Println("Output: ", fourSum1(nums, target))
}
