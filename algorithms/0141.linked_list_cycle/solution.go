package problems

type ListNode struct {
	Val  int
	Next *ListNode
}

// 141. Linked List Cycle (Easy)
// https://leetcode.com/problems/linked-list-cycle

// Double pointer solution
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slowHead := head
	fastHead := head

	for fastHead != nil && fastHead.Next != nil {
		slowHead = slowHead.Next
		fastHead = fastHead.Next.Next

		if fastHead == slowHead {
			return true
		}
	}

	return false
}
