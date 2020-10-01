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

func main() {
	s := "abcd"
	t := "abcde"
	fmt.Println("Output: ", string(findTheDifference(s, t)))
}
