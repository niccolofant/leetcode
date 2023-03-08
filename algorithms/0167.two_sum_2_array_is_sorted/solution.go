package algorithms

// 167. Two Sum II - Array is Sorted (Medium)
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

// Two pointers solution
func TwoSum(numbers []int, target int) []int {
	var result []int

	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum > target {
			right--
		} else {
			if sum < target {
				left++
			} else {
				result = []int{left + 1, right + 1}
				break
			}
		}
	}

	return result
}
