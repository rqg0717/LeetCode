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

func findTheDifference3(s string, t string) byte {
	result := byte(0)
	m := make(map[byte]int)
	for ch := range t {
		m[t[ch]]++
	}
	for ch := range s {
		m[s[ch]]--
	}
	for ch, count := range m {
		if count == 1 {
			result = ch
		}
	}
	return result
}

func main() {
	s := "abcd"
	t := "abcde"
	fmt.Println("Output: ", string(findTheDifference3(s, t)))
}
