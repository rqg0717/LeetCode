package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	dp := make([]int, n)
	ret := 1
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ret = max(ret, dp[i])
	}
	return ret
}

// MinUint is the minimum unsigned integer
const MinUint = 0

// MaxUint is the maximum unsigned integer
const MaxUint = ^uint(MinUint)

// MaxInt is the maximum signed integer
const MaxInt = int(MaxUint >> 1)

// MinInt is the minimum signed integer
const MinInt = -MaxInt - 1

func lengthOfLIS1(nums []int) int {
	dp := []int{MinInt}
	for _, val := range nums {
		if val > dp[len(dp)-1] {
			dp = append(dp, val)
		} else {
			left, right := 0, len(dp)
			for left < right {
				mid := left + (right-left)>>1
				if dp[mid] < val {
					left = mid + 1
				} else {
					right = mid
				}
			}
			dp[right] = val
		}
	}
	return len(dp) - 1
}

func main() {
	nums := []int{10, 9, 4, 5, 2, 7, 1, 0, 6}
	fmt.Println("Output: ", lengthOfLIS1(nums))
}
