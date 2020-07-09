package main

import "fmt"

func fib(N int) int {
	/* fastest, it did say 0 ≤ N ≤ 30 lol */
	// fibArray := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34,
	// 	55, 89, 144, 233, 377, 610,
	// 	987, 1597, 2584, 4181, 6765, 10946,
	// 	17711, 28657, 46368, 75025, 121393,
	// 	196418, 317811, 514229, 832040}

	// return fibArray[N]

	if N <= 1 {
		return N
	}

	/* slowest, O(2^N) and O(N)
	return fib(N-1) + fib(N-2)
	*/

	/* from top to bottom O(N) and O(N)
	return storage(N)
	*/

	// from bottom to top O(N) and O(1)
	if N == 2 {
		return 1
	}

	sum := 0
	previous := 1
	current := 1

	for i := 3; i <= N; i++ {
		sum = previous + current
		previous = current
		current = sum
	}
	return current
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
