package main

import (
	"fmt"
)

func findComplement(num int) int {
	mask := 1
	for mask <= num {
		mask <<= 1
	}
	mask = (mask - 1) ^ num
	return mask
}

func findComplement1(num int) int {
	mask := 1
	for mask < num {
		mask = (mask << 1) + 1
	}
	mask ^= num
	return mask
}

func main() {
	num := 5
	fmt.Println("Output: ", findComplement(num))
}
