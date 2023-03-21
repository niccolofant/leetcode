# Invert Binary Tree

## Problem

Given the `root` of a binary tree, invert the tree, and return its root.

## Examples

![Example1](https://assets.leetcode.com/uploads/2021/03/14/invert1-tree.jpg)

```text
Input: root = [4,2,7,1,3,6,9]

Output: [4,7,2,9,6,3,1]
```

![Example2](https://assets.leetcode.com/uploads/2021/03/14/invert2-tree.jpg)

```text
Input: root = [2,1,3]

Output: [2,3,1]
```

```text
Input: root = []

Output: []
```

## Solution

### Recursive Approach

if the current node is `nil` we return it.
Otherwise, we swap its `left` child with its `right` child and recursively
repeat the process with its `left` child and with its `right` child.

#### Code

##### Go

```go
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
```

#### Complexity Analysis

- **Time Complexity** : `O(n)`, where `n` is the number of nodes in the tree.

- **Space Complexity** : `O(n)`. The number of recursive calls is bound by the height of the tree. In the worst case, the tree is linear and the height is in `O(n)`. Therefore, space complexity due to recursive calls on the stack is `O(n)` in the worst case.
