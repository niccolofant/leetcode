#  Binary Search Tree Iterator

## Problem

Implement the `BSTIterator` class that represents an iterator over the **in-order traversal** of a binary search tree (BST):

- `BSTIterator(TreeNode root)` Initializes an object of the `BSTIterator class`. The `root` of the BST is given as part of the constructor. The pointer should be initialized to a non-existent number smaller than any element in the BST.

- `boolean hasNext()` Returns `true` if there exists a number in the traversal to the right of the pointer, otherwise returns `false`.

- `int next()` Moves the pointer to the right, then returns the number at the pointer.

Notice that by initializing the pointer to a non-existent smallest number, the first call to next() will return the smallest element in the BST.

You may assume that `next()` calls will always be valid. That is, there will be at least a next number in the in-order traversal when `next()` is called.

## Example

![BST Tree](https://assets.leetcode.com/uploads/2018/12/25/bst-tree.png)

```text
Input:
["BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"]
[[[7, 3, 15, null, null, 9, 20]], [], [], [], [], [], [], [], [], []]
Output
[null, 3, 7, true, 9, true, 15, true, 20, false]

Explanation:
BSTIterator bSTIterator = new BSTIterator([7, 3, 15, null, null, 9, 20]);
bSTIterator.next();    // return 3
bSTIterator.next();    // return 7
bSTIterator.hasNext(); // return True
bSTIterator.next();    // return 9
bSTIterator.hasNext(); // return True
bSTIterator.next();    // return 15
bSTIterator.hasNext(); // return True
bSTIterator.next();    // return 20
bSTIterator.hasNext(); // return False
```

## Solution

### Pre-build In Order Traversal

The pre-build in order traversal solution build the in order traversal of the binary search tree and stores it in a slice.
The `BSTIterator` class store the `InOrderTraversal` slice and the `CurrentIndex` in order to keep track of the node
(Note that the slice + index can be replaced by a simple stack).
The `next()` method simply increment the `CurrentIndex` and return the value at the `CurrentIndex`.
The `hasNext()` method simply check if the incremented `CurrentIndex` is smaller than the length of the `InOrderTraversal` slice.

#### Code

##### Go

```go
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
```

##### Complexity Analysis

- **Time Complexity** : `O(n)` where `n` is the number of nodes in the binary search tree.

- **Space Complexity** : `O(n)` where `n` is the number of nodes in the binary search tree.
