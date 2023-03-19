package algorithms

// 509. Fibonacci Number (Easy)
// https://leetcode.com/problems/fibonacci-number

// Memoization solution
func Fib(n int) int {
	return FibMemo(n, make(map[int]int))
}

func FibMemo(n int, memo map[int]int) int {
	if memoizedValue, ok := memo[n]; ok {
		return memoizedValue
	}

	if n == 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	memo[n] = FibMemo(n-1, memo) + FibMemo(n-2, memo)

	return memo[n]
}
