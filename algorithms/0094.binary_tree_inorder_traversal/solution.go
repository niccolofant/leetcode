package algorithms

import (
	data_structures "github.com/ocintnaf/leetcode/data_structures/stack"
)

// Definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 94. Binary Tree Inorder Traversal (Easy)
// https://leetcode.com/problems/binary-tree-inorder-traversal/description/

// Iterative solution
func InorderTraversal(root *TreeNode) []int {
	var result []int
	nodesStack := new(data_structures.Stack[*TreeNode])

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

// Recursive solution
func InorderTraversalRecursive(root *TreeNode) []int {
	var result []int

	if root != nil {
		result = append(result, InorderTraversalRecursive(root.Left)...)
		result = append(result, root.Val)
		result = append(result, InorderTraversalRecursive(root.Right)...)
	}

	return result
}
