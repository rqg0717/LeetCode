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

/**
 *	O( log(n)+log(n-1)+log(n-2)+...+log(1) ) => O( log(1⋅...⋅(n-2)⋅(n-1)⋅n) ) => O(log(n!))
 */
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

func searchMatrix1(matrix [][]int, target int) bool {
	// start with the bottom left
	row := len(matrix) - 1
	column := 0
	for row >= 0 && column <= (len(matrix[0])-1) {
		if matrix[row][column] > target {
			row--
		} else if matrix[row][column] < target {
			column++
		} else {
			return true
		}
	}
	return false
}

func main() {
	target := 27
	matrix := [][]int{
		{1, 4, 7, 11, 15, 17},
		{2, 5, 8, 12, 19, 20},
		{3, 6, 9, 16, 22, 23},
		{10, 13, 14, 17, 24, 25},
		{18, 21, 23, 26, 28, 30}}

	fmt.Println("input Matrix: ", matrix)
	fmt.Println("Output: ", searchMatrix1(matrix, target))
}
