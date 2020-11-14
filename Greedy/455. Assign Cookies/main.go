package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
		}
		j++
	}
	return i
}

func main() {
	g := []int{1, 2}
	s := []int{1, 2, 3}
	fmt.Println("Result: ", findContentChildren(g, s))
}
