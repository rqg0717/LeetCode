package main

import (
	"fmt"
	"math"
)

// ArrayReader contains int array
type ArrayReader struct {
	nums []int
}

func (ar *ArrayReader) get(index int) int {
	if index < 0 || index >= len(ar.nums) {
		return math.MaxInt32
	}
	return ar.nums[index]
}

func search(reader ArrayReader, target int) int {
	left, right := 0, math.MaxInt32
	for left < right {
		mid := left + (right-left)>>1
		x := reader.get(mid)
		if x == target {
			return mid
		} else if x < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return -1
}

func main() {
	target := 9
	reader := ArrayReader{[]int{-1, 0, 3, 5, 9, 12}}
	fmt.Println("Output: ", search(reader, target))
}
