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

func main() {
	s := "ababb"
	p := "ab"
	fmt.Println("Output: ", findAnagrams(s, p))
}
