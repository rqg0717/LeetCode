package main

import (
	"fmt"
)

func reverseBits(num uint32) uint32 {
	result := uint32(0)
	i := 0
	for i < 32 {
		result = (result << 1) + (num & 1)
		num >>= 1
		i++
	}
	return result
}

func reverseBits1(num uint32) uint32 {
	num = (num >> 16) | (num << 16)
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
	return num
}

func main() {
	num := uint32(0b00000010100101000001111010011100)
	fmt.Println("Output: ", reverseBits(num))
}
