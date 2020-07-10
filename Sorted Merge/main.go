package main

import (
	"fmt"
	"sort"
)

func merge(A []int, m int, B []int, n int) {
	copy(A[m:], B[:])
	sort.Ints(A)
}

func merge1(A []int, m int, B []int, n int) {
	indexA := m - 1
	indexB := n - 1
	index := len(A) - 1
	for indexA >= 0 && indexB >= 0 {
		if A[indexA] > B[indexB] {
			A[index] = A[indexA]
			indexA--
		} else {
			A[index] = B[indexB]
			indexB--
		}
		index--
	}
	for indexB >= 0 {
		A[index] = B[indexB]
		indexB--
		index--
	}
}

func main() {
	A := []int{1, 2, 3, 0, 0, 0}
	m := 3
	B := []int{2, 5, 6}
	n := 3
	merge1(A, m, B, n)
	fmt.Println("Output: ", A)
}
