package algorithms

type ListNode struct {
	Val  int
	Next *ListNode
}

// 142. Linked List Cycle II (Medium)
// https://leetcode.com/problems/linked-list-cycle-ii

// Hash table solution
func DetectCycle(head *ListNode) *ListNode {
	nodesMap := make(map[*ListNode]*ListNode)

	currentNode := head

	for currentNode != nil {
		if node, ok := nodesMap[currentNode]; ok {
			return node
		}

		nodesMap[currentNode] = currentNode
		currentNode = currentNode.Next
	}

	return nil
}
