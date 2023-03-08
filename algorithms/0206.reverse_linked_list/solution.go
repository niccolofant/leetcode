package problems

type ListNode struct {
	Val  int
	Next *ListNode
}

// 206. Reverse Linked List (Easy)
// https://leetcode.com/problems/reverse-linked-list/

// Iterative solution
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
