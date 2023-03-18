package algorithms

// 704. Binary Search (Easy)
// https://leetcode.com/problems/binary-search

// Recursive solution
func SearchRecursive(nums []int, target int) int {
	return BinarySearchRecursive(nums, 0, len(nums)-1, target)
}

func BinarySearchRecursive(nums []int, low int, high int, target int) int {
	if low <= high {
		mid := (low + high) / 2

		if target == nums[mid] {
			return mid
		} else {
			if target > nums[mid] {
				return BinarySearchRecursive(nums, mid+1, high, target)
			} else {
				return BinarySearchRecursive(nums, low, mid-1, target)
			}
		}
	}

	return -1
}

// Iterative solution
func SearchIterative(nums []int, target int) int {
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

	return -1
}
