package main

import "fmt"

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// dynamic programming
func findMaxForm(strs []string, m int, n int) int {
	l := len(strs)
	if l == 0 {
		return 0
	}

	/* count numbers of '0' and '1' */
	count := make([][]int, 2)
	count[0] = make([]int, l)
	count[1] = make([]int, l)
	for i := 0; i < l; i++ {
		strl := len(strs[i])
		for j := 0; j < strl; j++ {
			r := strs[i][j] - '0'
			count[r][i]++
		}
	}
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= l; i++ {
		for j := m; j >= count[0][i-1]; j-- {
			for k := n; k >= count[1][i-1]; k-- {
				dp[j][k] = max(dp[j][k], dp[j-count[0][i-1]][k-count[1][i-1]]+1)
			}
		}
	}
	return dp[m][n]
}

func main() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m := 5
	n := 3
	fmt.Println("total strings can be formed: ", findMaxForm(strs, m, n))
}
