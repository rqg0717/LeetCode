package main

import "fmt"

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

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4, 4, -2}
	target := 0
	fmt.Println("Output: ", fourSum(nums, target))
}
