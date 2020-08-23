package main

import (
	"fmt"
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func removeDuplicateLetters(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	for i := 0; i < len(r)-1; {
		if r[i] == r[i+1] {
			r = append(r[:i], r[i+1:]...)
		} else {
			i++
		}
	}
	return string(r)
}

func removeDuplicateLetters1(s string) string {
	n := len(s)
	m1 := make(map[byte]int)
	for i := 0; i < n; i++ {
		m1[s[i]] = i
	}
	m2 := make(map[byte]bool)
	result := ""
	for i := 0; i < n; i++ {
		if m2[s[i]] {
			continue
		}
		for len(result) != 0 && result[len(result)-1] > s[i] && m1[result[len(result)-1]] > i {
			m2[result[len(result)-1]] = false
			result = result[:len(result)-1]
		}
		m2[s[i]] = true
		result += string(s[i])
	}
	return result
}

func main() {
	s := "cbacdcbc"
	fmt.Println("Output: ", removeDuplicateLetters(s))
}
