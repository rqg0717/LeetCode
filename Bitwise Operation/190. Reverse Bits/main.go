package main

import (
	"fmt"
)

func reverseBits(num uint32) uint32 {
	result := uint32(0)
	for num > 0 {
		result = (result << 1) + (num & 1)
		num >>= 1
	}
	return result
}

func main() {
	num := uint32(0b11111111111111111111111111111101)
	fmt.Println("Output: ", reverseBits(num))
}
