package main

import "fmt"

/* global variable declaration */
var pick int

func guess(num int) int {
	if num == pick {
		return 0
	} else if num < pick {
		return -1
	} else {
		return 1
	}
}

func guessNumber(n int) int {
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)>>1
		switch ans := guess(mid); ans {
		case 0:
			return mid
		case -1:
			left = mid + 1
		case 1:
			right = mid - 1
		}
	}
	return -1
}

func main() {
	pick = 6
	n := 20
	fmt.Println("Output: ", guessNumber(n))
}
