package algorithms

// 268. Missing Number (Easy)
// https://leetcode.com/problems/missing-number/
func MissingNumber(nums []int) int {
	n := len(nums)
	expectedSumOfNums := (n * (n + 1)) / 2
	actualSumOfNums := 0

	for _, num := range nums {
		actualSumOfNums += num
	}

	return expectedSumOfNums - actualSumOfNums
}
