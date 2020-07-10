package main

import (
	"fmt"
	"sort"
)

func merge(A []int, m int, B []int, n int) {
	copy(A[m:], B[:])
	sort.Ints(A)
}

func main() {
	A := []int{1, 2, 3, 0, 0, 0}
	m := 3
	B := []int{2, 5, 6}
	n := 3
	merge(A, m, B, n)
	fmt.Println("Output: ", A)
}
