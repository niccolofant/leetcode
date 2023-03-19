package algorithms

// 35. Search Insert Position (Easy)
// https://leetcode.com/problems/search-insert-position

// Binary Search Solution
func SearchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2

		if target == nums[mid] {
			return mid
		} else {
			if target > nums[mid] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return low
}
