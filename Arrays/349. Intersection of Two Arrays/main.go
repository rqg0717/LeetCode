package main

import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {
	m1 := map[int]int{}
	m2 := map[int]int{}
	results := []int{}
	for _, value := range nums1 {
		m1[value] = value
	}

	for _, value := range nums2 {
		m2[value] = value
	}

	for _, value := range m1 {
		if value, ok := m2[value]; ok {
			results = append(results, value)
		}
	}

	return results
}

func intersection1(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	results := []int{}
	for _, value := range nums1 {
		m[value]++
	}
	for _, value := range nums2 {
		if m[value] > 0 {
			m[value] = 0
			results = append(results, value)
		}
	}
	return results
}

func intersection2(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for _, value := range nums1 {
		m[value]++
	}
	index := 0
	for _, value := range nums2 {
		if m[value] > 0 {
			m[value] = 0
			nums2[index] = value
			index++
		}
	}
	return nums2[0:index]
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println("Output: ", intersection2(nums1, nums2))
}
