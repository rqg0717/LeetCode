package main

import (
	"fmt"
	"sort"
)

type interval [][]int

func (i interval) Len() int {
	return len(i)
}

func (i interval) Less(x, y int) bool {
	if i[x][1] != i[y][1] {
		return i[x][1] < i[y][1]
	}
	return i[x][0] < i[y][0]
}

func (i interval) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sort.Sort(interval(intervals))
	previous, result := 0, 1
	for i := 1; i < n; i++ {
		if intervals[i][0] >= intervals[previous][1] {
			result++
			previous = i
		}
	}
	return n - result
}

func main() {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	fmt.Println("Output: ", eraseOverlapIntervals(intervals))
}
