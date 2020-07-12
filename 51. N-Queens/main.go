package main

import "fmt"

func solveNQueens(n int) [][]string {
	// there are solutions iff n == 1 and n >= 4
	if n < 4 && n != 1 {
		return nil
	}

	results := make([][]int, 0)
	// Depth-First-Search
	dfs([]int{}, n, 0, 0, 0, &results)
	// print results
	return generateResults(results, n)
}

func dfs(rows []int, n int, columns, leftright, rightleft int, results *[][]int) {
	// end condition
	if len(rows) == n {
		sub := make([]int, n)
		copy(sub, rows)
		(*results) = append((*results), sub)
		return
	}

	/*
		a := byte(0b11110000)
		(^a) = 0b00001111
		-a = 0b00010000
		(a & -a) = 0b00010000
	*/
	a := (^(columns | leftright | rightleft)) & ((1 << uint(n)) - 1) // currently available columns
	for a != 0 {
		i := a & (-a)
		a ^= i
		cols := 0
		j := (1 << uint(n-1))
		for i&j == 0 {
			cols++
			j >>= 1
		}
		// moves to next row
		// leftright [0, 1] to [1, 2]
		// rightleft [0, 2] to [1, 1]
		dfs(append(rows, cols), n, columns^i, (leftright^i)>>1, (rightleft^i)<<1, results)
	}
}

func generateResults(results [][]int, n int) (res [][]string) {
	for _, r := range results {
		var str []string
		for _, val := range r {
			char := ""
			for position := 0; position < n; position++ {
				if position == val {
					char += "Q"
				} else {
					char += "."
				}
			}
			str = append(str, char)
		}
		res = append(res, str)
	}
	return res
}

func main() {
	n := 8
	fmt.Println("Output:", solveNQueens(n))
}
