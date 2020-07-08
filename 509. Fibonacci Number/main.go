package main

import "fmt"

func fib(N int) int {
	/* fastest, it did say 0 ≤ N ≤ 30 lol */
	fibArray := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34,
		55, 89, 144, 233, 377, 610,
		987, 1597, 2584, 4181, 6765, 10946,
		17711, 28657, 46368, 75025, 121393,
		196418, 317811, 514229, 832040}

	return fibArray[N]

	// if N <= 1 {
	// 	return N
	// }

	/* slowest, O(2^N) and O(N)
	return fib(N-1) + fib(N-2)
	*/

	/* O(N) and O(N)
	return storage(N)
	*/

	// O(N) and O(1)
	/*if N == 2 {
		return 1
	}

	current := 0
	previous1 := 1
	previous2 := 1

	for i := 3; i <= N; i++ {
		current = previous1 + previous2
		previous2 = previous1
		previous1 = current
	}
	return current
	*/
}

func storage(N int) int {

	cache := map[int]int{0: 0, 1: 1}

	for i := 2; i <= N; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}
	return cache[N]
}

func main() {
	for N := 0; N <= 30; N++ {
		fmt.Println("output: ", N, fib(N))
	}
}
