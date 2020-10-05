package main

import (
	"fmt"
)

func countPrimeSetBits(L int, R int) int {
	primes := []uint{2, 3, 5, 7, 11, 13, 17, 19} // 10^6
	mask := 0
	for _, val := range primes {
		mask |= 1 << val
	}
	sum := 0
	for i := L; i <= R; i++ {
		n := i
		count := 0
		for n > 0 {
			count++
			n = n & (n - 1)
		}
		if mask&(1<<count) > 0 {
			sum++
		}
	}
	return sum
}

func main() {
	L := 6
	R := 10
	fmt.Println("Output: ", countPrimeSetBits(L, R))
}
