package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	for i, v := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			if target-v == numbers[j] {
				return []int{i + 1, j + 1}
			}
		}
	}
	return nil
}

func twoSum1(numbers []int, target int) []int {
	n := len(numbers)
	for i := 0; i < n; i++ {
		left, right := i+1, n-1
		for left <= right {
			mid := left + (right-left)>>1
			if numbers[mid] == target-numbers[i] {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] > target-numbers[i] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return []int{-1, -1}
}

func twoSum2(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{-1, -1}
}

func main() {
	numbers := []int{1, 2, 3, 5, 7, 11}
	target := 9
	fmt.Println("Output: ", twoSum1(numbers, target))
}
