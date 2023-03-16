# Convert Sorted Array to Binary Search Tree

## Problem

Given an integer array `nums` where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.
A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.

## Example

![Example1](https://assets.leetcode.com/uploads/2021/02/18/btree1.jpg)
![Example2](https://assets.leetcode.com/uploads/2021/02/18/btree2.jpg)

```text
Input: nums = [-10,-3,0,5,9]

Output: [0,-3,9,-10,null,5] or [0,-10,5,null,-3,null,9]
```

## Solution

### Recursive Solution

#### Code

##### Go

```go
func SortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
    return nil
    }

    mid := len(nums) / 2
    root := &TreeNode{Val: nums[mid]}
    root.Left = SortedArrayToBST(nums[:mid])
    root.Right = SortedArrayToBST(nums[mid+1:])

    return root
}
```

#### Complexity Analysis

- **Time Complexity** : `O(n)`, where `n` is the number of nodes in the binary search tree.

- **Space Complexity** : `O(n)`, where `n` is the number of nodes in the binary search tree.
