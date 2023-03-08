// Definition for a binary tree node.
class TreeNode {
  val: number
  left: TreeNode | null
  right: TreeNode | null
  constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
    this.val = val === undefined ? 0 : val
    this.left = left === undefined ? null : left
    this.right = right === undefined ? null : right
  }
}

// 94. Binary Tree Inorder Traversal (Easy)
// https://leetcode.com/problems/binary-tree-inorder-traversal/description/

// Iterative solution
const inorderTraversal = (root: TreeNode | null): number[] => {
  const result: number[] = []
  const nodesStack: TreeNode[] = []

  let currentNode: TreeNode | null = root

  while (currentNode || nodesStack.length) {
    if (currentNode) {
      nodesStack.push(currentNode)
      currentNode = currentNode.left
    } else {
      currentNode = nodesStack.pop()

      result.push(currentNode.val)

      currentNode = currentNode.right
    }
  }

  return result
}
