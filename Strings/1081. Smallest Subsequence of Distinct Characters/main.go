package main

import (
	"fmt"
)

func smallestSubsequence(text string) string {
	letters := [26]int{}
	for _, val := range text {
		letters[val-'a']++
	}
	r := make([]rune, 0)
	visited := [26]bool{}
	for _, ch := range text {
		letters[ch-'a']--
		if visited[ch-'a'] {
			continue
		}
		for len(r) > 0 && ch < r[len(r)-1] && letters[r[len(r)-1]-'a'] > 0 {
			visited[r[len(r)-1]-'a'] = false
			r = r[:len(r)-1]
		}
		r = append(r, ch)
		visited[ch-'a'] = true
	}

	return string(r)
}

func main() {
	s := "cbacdcbc"
	fmt.Println("Output: ", smallestSubsequence(s))
}
