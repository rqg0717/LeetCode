package main

import (
	"fmt"
)

func validUtf8(data []int) bool {
	bit := 0
	for i := 0; i < len(data); i++ {
		mask1, mask2 := 1<<7, 1<<6
		if bit != 0 {
			if !((data[i]&mask1 != 0) && (data[i]&mask2) == 0) {

				return false
			}
			bit--
			continue
		}
		for data[i]&mask1 != 0 {
			bit++
			mask1 >>= 1
		}
		if bit == 0 {
			continue
		}
		if bit == 1 || bit > 4 {
			return false
		}
		bit--
	}
	return bit == 0
}

func main() {
	data := []int{235, 140, 4}
	fmt.Println("Output: ", validUtf8(data))
}
