package algorithms

// Definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 173. Binary Search Tree Iterator (Medium)
// https://leetcode.com/problems/binary-search-tree-iterator/description/

type BSTIterator struct {
	InOrderTraversal []int
	CurrentIndex     int
}

func Constructor(root *TreeNode) BSTIterator {
	inOrderTraversal := InOrderTraversal(root)

	bstIterator := new(BSTIterator)
	bstIterator.InOrderTraversal = inOrderTraversal
	bstIterator.CurrentIndex = -1

	return *bstIterator
}

func (bstIterator *BSTIterator) Next() int {
	bstIterator.CurrentIndex += 1
	return bstIterator.InOrderTraversal[bstIterator.CurrentIndex]
}

func (bstIterator *BSTIterator) HasNext() bool {
	return bstIterator.CurrentIndex+1 < len(bstIterator.InOrderTraversal)
}

func InOrderTraversal(root *TreeNode) []int {
	var result []int

	if root != nil {
		result = append(result, InOrderTraversal(root.Left)...)
		result = append(result, root.Val)
		result = append(result, InOrderTraversal(root.Right)...)
	}

	return result
}
