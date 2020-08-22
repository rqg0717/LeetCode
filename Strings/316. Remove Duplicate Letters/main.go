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
	for i := 0; i+1 < len(r); {
		if r[i] == r[i+1] {
			r = append(r[:i], r[i+1:]...)
		} else {
			i++
		}
	}
	return string(r)
}

func main() {
	s := "bcabc"
	fmt.Println("Output: ", removeDuplicateLetters(s))
}
