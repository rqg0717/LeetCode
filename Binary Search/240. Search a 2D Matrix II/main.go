package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func binarySearch(matrix [][]int, target int, start int, isVertical bool) bool {
	left := start
	right := 0
	if isVertical {
		right = len(matrix[0]) - 1
	} else {
		right = len(matrix) - 1
	}
	for right >= left {
		mid := left + (right-left)>>2
		if isVertical {
			if matrix[start][mid] < target {
				left = mid + 1
			} else if matrix[start][mid] > target {
				right = mid - 1
			} else {
				return true
			}
		} else {
			if matrix[mid][start] < target {
				left = mid + 1
			} else if matrix[mid][start] > target {
				right = mid - 1
			} else {
				return true
			}
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	shorter := min(len(matrix), len(matrix[0]))
	isFound := false
	for i := 0; i < shorter; i++ {
		isFound = binarySearch(matrix, target, i, true)
		if isFound {
			break
		}
		isFound = binarySearch(matrix, target, i, false)
		if isFound {
			break
		}
	}
	return isFound
}

func main() {
	target := 16
	matrix := [][]int{
		{1, 4, 7, 11, 15, 17},
		{2, 5, 8, 12, 19, 20},
		{3, 6, 9, 16, 22, 23},
		{10, 13, 14, 17, 24, 25},
		{18, 21, 23, 26, 28, 30}}

	fmt.Println("input Matrix: ", matrix)
	fmt.Println("Output: ", searchMatrix(matrix, target))
}
