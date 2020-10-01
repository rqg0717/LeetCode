package main

import "fmt"

func findTheDifference(s string, t string) byte {
	result := byte(0)
	for _, ch := range s {
		result = result ^ byte(ch)
	}
	for _, ch := range t {
		result = result ^ byte(ch)
	}
	return result
}

func findTheDifference1(s string, t string) byte {
	result := byte(0)
	for i := 0; i < len(s); i++ {
		result ^= s[i] ^ t[i]
	}
	result ^= t[len(s)]
	return result
}

func findTheDifference2(s string, t string) byte {
	result := 0
	for _, ch := range t {
		result += int(ch)
	}
	for _, ch := range s {
		result -= int(ch)
	}
	return byte(result)
}

func main() {
	s := "abcd"
	t := "abcde"
	fmt.Println("Output: ", string(findTheDifference1(s, t)))
}
