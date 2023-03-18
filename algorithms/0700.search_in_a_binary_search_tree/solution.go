package algorithms

// Definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 700. Search in a Binary Search Tree (Easy)
// https://leetcode.com/problems/search-in-a-binary-search-tree

// Recursive solution
func SearchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if val == root.Val {
		return root
	} else {
		if val > root.Val {
			return SearchBST(root.Right, val)
		} else {
			return SearchBST(root.Left, val)
		}
	}
}
