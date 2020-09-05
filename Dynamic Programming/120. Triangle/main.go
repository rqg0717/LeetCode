package main

import (
	"fmt"
)

func dfs(i, j, min, sum int, triangle [][]int) int {
	if i == len(triangle) {
		if min > sum {
			min = sum
		}
		return min
	}
	a := dfs(i+1, j, min, sum+triangle[i][j], triangle)
	b := dfs(i+1, j+1, min, sum+triangle[i][j], triangle)
	if a < b {
		return a
	}
	return b
}

func minimumTotal(triangle [][]int) int {
	i, j, min, sum := 0, 0, 0, 0
	for n := 0; n < len(triangle); n++ {
		min = min + triangle[n][0]
	}
	return dfs(i, j, min, sum, triangle)
}

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println("Output: ", minimumTotal(triangle))
}
