# Same Tree

## Problem

Given the roots of two binary trees `p` and `q`, write a function to check if they are the same or not.
Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

## Example

![Example1](https://assets.leetcode.com/uploads/2020/12/20/ex1.jpg)

```text
Input: p = [1,2,3], q = [1,2,3]

Output: true
```

![Example2](https://assets.leetcode.com/uploads/2020/12/20/ex2.jpg)

```text
Input: p = [1,2], q = [1,null,2]

Output: false
```

## Solution

### Recursive Solution

The recursive solution is pretty straightforward. We check if both nodes are `nil` and return `true` if they are.
If both nodes are not `nil`, we check if their values are equal and if the left and right subtrees are equal.
Otherwise, we return `false`.

#### Code

##### Go

```go
func IsSameTreeRecursive(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }

    if p != nil && q != nil {
        return p.Val == q.Val && IsSameTreeRecursive(p.Left, q.Left) && IsSameTreeRecursive(p.Right, q.Right)
    }

    return false
}
```

#### Complexity Analysis

- **Time Complexity** : `O(n)`, where `n` is the number of nodes in the tree. We have to visit each node at least once.

- **Space Complexity** : `O(n)`. The number of recursive calls is bound by the height of the tree. In the worst case, the tree is linear and the height is in `O(n)`. Therefore, space complexity due to recursive calls on the stack is `O(n)` in the worst case.
