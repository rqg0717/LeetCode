package main

import "fmt"

func reverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func reverseString1(s []byte) {
	n := len(s)
	mid := n / 2
	for i := 0; i < mid; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
}

func main() {
	s := []byte("Hello World!")
	reverseString1(s)
	fmt.Println("Output: ", string(s))
}
