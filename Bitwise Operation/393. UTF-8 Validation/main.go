package main

import (
	"fmt"
)

func validUtf8(data []int) bool {
	n := len(data)
	if n == 0 {
		return false
	}
	bit := 0
	for i := 0; i < n; i++ {
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

func length(num int) int {
	// 0xxxxxxx
	if num>>7 == 0 {
		return 0
	}
	// 110xxxxx 10xxxxxx
	if num>>5 == 6 {

		return 1
	}
	// 1110xxxx 10xxxxxx 10xxxxxx
	if num>>4 == 14 {
		return 2
	}
	// 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
	if num>>3 == 30 {
		return 3
	}
	return -1
}

func validUtf8a(data []int) bool {
	n := len(data)
	if n == 0 {
		return false
	}
	l := 0
	for i := 0; i < n; i++ {
		if l == 0 {
			l = length(data[i])
		} else {
			// 10xxxxxx
			if data[i]&(3<<6) != 1<<7 {
				return false
			}
			l--
		}
		if l == -1 {
			return false
		}
	}
	return l == 0

}

func main() {
	data := []int{235, 140, 4}
	fmt.Println("Output: ", validUtf8a(data))
}
