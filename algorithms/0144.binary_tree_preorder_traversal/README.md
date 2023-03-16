# Binary Tree Preorder Traversal

## Problem

Given the `root` of a binary tree, return the preorder traversal of its nodes' values.
The preorder traversal of a binary tree is the preorder traversal of its root node followed by the preorder traversal of its left subtree followed by the preorder traversal of its right subtree.

## Example

![Binary Tree](https://assets.leetcode.com/uploads/2020/09/15/inorder_1.jpg)

```text
Input: root = [1,null,2,3]

Output: [1,2,3]
```

## Solution

### Recursive Approach

The recursive approach uses recursion to traverse the binary tree.
First, the value of the root node is added to the result.
Then, the recursive approach is called on the left subtree.
Finally, the recursive approach is called on the right subtree.
When the left and right subtrees are `nil`, the recursion ends and the result is returned.

#### Code

##### Go

```go
func PreorderTraversalRecursive(root *TreeNode) []int {
    result := []int{}

    if root != nil {
        result = append(result, root.Val)
        result = append(result, PreorderTraversalRecursive(root.Left)...)
        result = append(result, PreorderTraversalRecursive(root.Right)...)
    }

    return result
}
```

#### Complexity Analysis

- **Time Complexity** : `O(n)`, where `n` is the number of nodes in the binary tree.

- **Space Complexity** : `O(n)`, where `n` is the number of nodes in the binary tree.

---

### Iterative Approach

The iterative approach uses a stack to traverse the binary tree.
First, we initialize a `currentNode` variable to point to the root node.
Then, while the `currentNode` is not `nil` or the stack is not empty, we do the following: - If the `currentNode` is not `nil`, then we add the value of the `currentNode` to the result, push its right child to the stack, and set the `currentNode` to its left child. - If the `currentNode` is `nil`, then we pop the top element from the stack and set the `currentNode` to it.
At the end, the result is returned.

#### Code

##### Go

```go
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
```

#### Complexity Analysis

- **Time Complexity** : `O(n)`, where `n` is the number of nodes in the binary tree.

- **Space Complexity** : `O(n)`, where `n` is the number of nodes in the binary tree.
