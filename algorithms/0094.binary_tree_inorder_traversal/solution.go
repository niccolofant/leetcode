package algorithms

// Definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 94. Binary Tree Inorder Traversal (Easy)
// https://leetcode.com/problems/binary-tree-inorder-traversal/description/

// Iterative solution
func inorderTraversal(root *TreeNode) []int {
	var result []int
	nodesStack := new(Stack)

	currentNode := root

	for currentNode != nil || !nodesStack.IsEmpty() {
		if currentNode != nil {
			nodesStack.Push(currentNode)
			currentNode = currentNode.Left
		} else {
			currentNode = nodesStack.Pop()

			result = append(result, currentNode.Val)

			currentNode = currentNode.Right
		}
	}

	return result
}
