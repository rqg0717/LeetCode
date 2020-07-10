package main

import "fmt"

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	results := [][]int{}

	for i, num := range nums {
		temp := make([]int, len(nums)-1)
		copy(temp[0:], nums[0:i])
		copy(temp[i:], nums[i+1:])

		sub := permute(temp)
		for _, j := range sub {
			results = append(results, append(j, num))
		}
	}
	return results
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println("Output: ", permute(nums))
}
