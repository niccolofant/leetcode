package algorithms

import data_structures "github.com/ocintnaf/leetcode/data_structures/stack"

// Definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 144. Binary Tree Preorder Traversal (Easy)
// https://leetcode.com/problems/binary-tree-preorder-traversal

// Recursive solution
func PreorderTraversalRecursive(root *TreeNode) []int {
	result := []int{}

	if root != nil {
		result = append(result, root.Val)
		result = append(result, PreorderTraversalRecursive(root.Left)...)
		result = append(result, PreorderTraversalRecursive(root.Right)...)
	}

	return result
}

// Iterative solution
func PreorderTraversalIterative(root *TreeNode) []int {
	result := []int{}
	stack := new(data_structures.Stack[*TreeNode])
	currentNode := root

	for currentNode != nil || !stack.IsEmpty() {
		if currentNode != nil {
			result = append(result, currentNode.Val)

			stack.Push(currentNode.Right)

			currentNode = currentNode.Left
		} else {
			currentNode = stack.Pop()
		}
	}

	return result
}
