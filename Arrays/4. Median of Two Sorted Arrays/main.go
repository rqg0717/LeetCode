package main

import (
	"fmt"
	"math"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	n1, n2 := len(nums1), len(nums2)
	left, right := 0, n1
	median1, median2 := 0, 0
	for left <= right {
		i := (left + right) / 2
		j := (n1+n2+1)/2 - i
		iMin := math.MinInt32
		if i != 0 {
			iMin = nums1[i-1]
		}
		iMax := math.MaxInt32
		if i != n1 {
			iMax = nums1[i]
		}
		jMin := math.MinInt32
		if j != 0 {
			jMin = nums2[j-1]
		}
		jMax := math.MaxInt32
		if j != n2 {
			jMax = nums2[j]
		}
		if iMin <= jMax {
			median1 = max(iMin, jMin)
			median2 = min(iMax, jMax)
			left = i + 1
		} else {
			right = i - 1
		}
	}
	if (n1+n2)%2 == 0 { // even length
		return float64(median1+median2) / 2.0
	}
	return float64(median1)
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	nums2 := []int{4, 5, 6}
	fmt.Println("Output: ", findMedianSortedArrays(nums1, nums2))
}
