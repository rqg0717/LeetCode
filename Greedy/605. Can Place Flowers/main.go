package main

import (
	"fmt"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	l := len(flowerbed)
	if n > l {
		return false
	}
	result := 0
	for i := 0; i < l; i++ {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == l-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			result++
		}
	}
	return result >= n
}

func main() {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 2
	fmt.Println("Output: ", canPlaceFlowers(flowerbed, n))
}
