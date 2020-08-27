package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row, column := len(matrix), len(matrix[0])
	left, right := 0, row*column-1
	for left+1 < right {
		mid := left + (right-left)>>1

		val := matrix[mid/column][mid%column]
		if val > target {
			right = mid
		} else if val < target {
			left = mid
		} else {
			fmt.Printf("location is [%d][%d].\n", mid/column, mid%column)
			return true
		}
	}
	if matrix[left/column][left%column] == target {
		fmt.Printf("location is [%d][%d].\n", left/column, left%column)
		return true
	}
	if matrix[right/column][right%column] == target {
		fmt.Printf("location is [%d][%d].\n", right/column, right%column)
		return true
	}
	return false
}

func main() {
	target := 0
	matrix := [][]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11}}

	fmt.Println("input Matrix: ", matrix)
	fmt.Println("Output: ", searchMatrix(matrix, target))
}
