package algorithms

// Definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 226. Invert Binary Tree (Easy)
// https://leetcode.com/problems/invert-binary-tree

// Recursive solution
func InvertTree(root *TreeNode) *TreeNode {
	if root != nil {
		temp := root.Left
		root.Left = root.Right
		root.Right = temp

		InvertTree(root.Left)
		InvertTree(root.Right)
	}

	return root
}
