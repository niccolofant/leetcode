package algorithms

// 2485. Find the Pivot Integer (Easy)
// https://leetcode.com/problems/find-the-pivot-integer/

// Add and subtract solution
func pivotInteger(n int) int {
	sumOfAllNumbers := n * (n + 1) / 2
	sumOfFirstKNumbers := 0

	k := 1

	for k <= n {
		sumOfFirstKNumbers += k

		if sumOfFirstKNumbers == sumOfAllNumbers {
			return k
		}

		sumOfAllNumbers -= k

		k++
	}

	return -1
}
