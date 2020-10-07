package main

import (
	"fmt"
)

func bitwiseComplement(N int) int {
	mask := 1
	for mask < N {
		mask = (mask << 1) + 1
	}
	mask ^= N
	return mask
}

func main() {
	N := 5
	fmt.Println("Output: ", bitwiseComplement(N))
}
