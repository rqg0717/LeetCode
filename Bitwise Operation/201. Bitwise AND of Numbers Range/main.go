package main

import (
	"fmt"
)

func rangeBitwiseAnd(m int, n int) int {
	shift := 0
	for m != n {
		m >>= 1
		n >>= 1
		shift++
	}
	return n << shift
}

func main() {
	m := 5
	n := 7
	fmt.Println("Output: ", rangeBitwiseAnd(m, n))
}
