package algorithms

// 27. Remove Element (Easy)
// https://leetcode.com/problems/remove-element

// Two Pointers solution
func RemoveElement(nums []int, val int) int {
	k := len(nums) - 1
	i := 0

	for i <= k {
		if nums[i] == val {
			nums[i] = -1 // Not really necessary
			nums[i], nums[k] = nums[k], nums[i]
			k--
		} else {
			i++
		}
	}

	return k + 1
}
