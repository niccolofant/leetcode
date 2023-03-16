package algorithms

// Definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 100. Same Tree (Easy)
// https://leetcode.com/problems/same-tree

// Recursive solution
func IsSameTreeRecursive(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q != nil {
		return p.Val == q.Val && IsSameTreeRecursive(p.Left, q.Left) && IsSameTreeRecursive(p.Right, q.Right)
	}

	return false
}
