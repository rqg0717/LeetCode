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
	for _, q := range queries {
		i, j, k := q[0], q[1], q[2]
		num := count[i] ^ count[j+1]
		ans := 0
		for num > 0 {
			ans += num & 1
			num >>= 1
		}
		results = append(results, ans <= (2*k+1))
	}
	return results
}

func canMakePaliQueries1(s string, queries [][]int) []bool {
	count := []int{0}
	for i, ch := range s {
		count = append(count, count[i]^(1<<(ch-'a')))
	}
	results := []bool{}
	for _, q := range queries {
		i, j, k := q[0], q[1], q[2]
		num := count[i] ^ count[j+1]
		num = (num & 0x55555555) + ((num >> 1) & 0x55555555)
		num = (num & 0x33333333) + ((num >> 2) & 0x33333333)
		num = (num & 0x0f0f0f0f) + ((num >> 4) & 0x0f0f0f0f)
		num = (num & 0x00ff00ff) + ((num >> 8) & 0x00ff00ff)
		num = (num & 0x0000ffff) + ((num >> 16) & 0x0000ffff)
		results = append(results, num <= (2*k+1))
	}
	return results
}

func main() {
	s := "abcda"
	queries := [][]int{{3, 3, 0}, {1, 2, 0}, {0, 3, 1}, {0, 3, 2}, {0, 4, 1}}
	fmt.Println("Output: ", canMakePaliQueries1(s, queries))
}
