package algorithms

type ListNode struct {
	Val  int
	Next *ListNode
}

// 876. Middle of the Linked List (Easy)
// https://leetcode.com/problems/middle-of-the-linked-list/

// Fast Pointer solution
func MiddleNode(head *ListNode) *ListNode {
	middleNode := head

	for head != nil && head.Next != nil {
		head = head.Next.Next
		middleNode = middleNode.Next
	}

	return middleNode
}
