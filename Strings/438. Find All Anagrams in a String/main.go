package main

import (
	"fmt"
)

func getProduct(str string) int {
	result := 1
	for i := 0; i < len(str); i++ {
		result *= int(str[i])
	}
	return result
}

func findAnagrams(s string, p string) []int {
	sn := len(s)
	pn := len(p)
	indexs := []int{}
	if sn == 0 || pn == 0 || pn > sn {
		return nil
	}
	pp := getProduct(p)
	for i := 0; i <= sn-pn; i++ {
		if getProduct(s[i:i+pn]) == pp {
			indexs = append(indexs, i)
		}
	}
	return indexs
}

func findAnagrams1(s string, p string) []int {
	sn := len(s)
	pn := len(p)
	indexs := []int{}
	if sn == 0 || pn == 0 || pn > sn {
		return nil
	}
	window := make(map[byte]int)
	pmap := make(map[byte]int)
	for i := 0; i < pn; i++ {
		pmap[p[i]]++
	}
	left, right := 0, 0
	valid := 0
	for right < sn {
		c1 := s[right]
		right++ // moves the window to the right
		if _, ok := pmap[c1]; ok {
			window[c1]++ // moves the char into the window
			if window[c1] == pmap[c1] {
				valid++
			}
		}
		for right-left >= pn {
			if valid == len(pmap) {
				indexs = append(indexs, left)
			}
			c2 := s[left]
			left++ // moves the window to the left
			if _, ok := window[c2]; ok {
				if window[c2] == pmap[c2] {
					valid--
				}
				window[c2]-- // moves the char out of the window
			}
		}
	}
	return indexs
}

func main() {
	s := "ababb"
	p := "ab"
	fmt.Println("Output: ", findAnagrams1(s, p))
}
