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
	return printResults(results, n)
}

func dfs(rows []int, n int, columns, leftright, rightleft int, results *[][]int) {
	// end condition for one solution
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
	a := (^(columns | leftright | rightleft)) & ((1 << uint(n)) - 1) // currently available columns marked as 1
	for a != 0 {
		i := a & (-a) // obtains the next available column
		a ^= i
		cols := 0
		j := (1 << uint(n-1))
		// checks every column
		for i&j == 0 {
			cols++
			j >>= 1 // moves to next column
		}
		// moves to next row
		dfs(append(rows, cols), n, columns^i, (leftright^i)>>1, (rightleft^i)<<1, results)
	}
}

func printResults(results [][]int, n int) (res [][]string) {
	for _, r := range results {
		str := []string{}
		for _, val := range r {
			char := ""
			for p := 0; p < n; p++ {
				if p == val {
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

func solveNQueens1(n int) [][]string {
	var res [][]string
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	columns, leftright, rightleft := 0, 0, 0
	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			var temp []string
			for _, v := range board {
				temp = append(temp, string(v))
			}
			res = append(res, temp)
			return
		}

		available := ((1 << n) - 1) & ^(columns | (leftright >> row) | (rightleft >> (n - 1 - row)))
		for available != 0 {
			p := available & -available
			available = available &^ p
			temp, col := p, 0
			for temp != 1 {
				temp, col = temp>>1, col+1
			}
			board[row][col], columns, leftright, rightleft = 'Q', columns|p, leftright|(p<<row), rightleft|(p<<(n-1-row))
			dfs(row + 1)
			board[row][col], columns, leftright, rightleft = '.', columns&^p, leftright&^(p<<row), rightleft&^(p<<(n-1-row))
		}
	}
	dfs(0)
	return res
}

func main() {
	n := 4
	fmt.Println("Output:", solveNQueens1(n))
}
