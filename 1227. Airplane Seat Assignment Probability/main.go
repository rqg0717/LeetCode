package main

// f(n) = 1/n(1+∑f(n−i+1)) = f(2)
func nthPersonGetsNthSeat(n int) float64 {
	if n == 1 {
		return 1
	}
	return 0.5
}

func main() {
	n := 100
	nthPersonGetsNthSeat(n)
}
