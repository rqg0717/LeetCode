package main

import (
	"fmt"
)

func canMakePaliQueries(s string, queries [][]int) []bool {
	count := []int{0}
	for i, ch := range s {
		count = append(count, count[i]^(1<<(ch-'a')))
	}
	results := []bool{}
	for _, v := range queries {
		i, j, k := v[0], v[1], v[2]
		num := (count[i] ^ count[j+1])
		ans := 0
		for num > 0 {
			ans += num & 1
			num >>= 1
		}
		results = append(results, ans/2 <= k)
	}
	return results
}

func main() {
	s := "abcda"
	queries := [][]int{{3, 3, 0}, {1, 2, 0}, {0, 3, 1}, {0, 3, 2}, {0, 4, 1}}
	fmt.Println("Output: ", canMakePaliQueries(s, queries))
}
