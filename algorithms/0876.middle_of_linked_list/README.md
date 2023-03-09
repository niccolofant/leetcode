#  Middle of the Linked List

## Problem

Given a non-empty, singly linked list with head node `head`, return a middle node of linked list.
If there are two middle nodes, return the second middle node.

## Example

```text
Input: [1,2,3,4,5]

Output: Node 3 from this list (Serialization: [3,4,5])
```

```text
Input: [1,2,3,4,5,6]

Output: Node 4 from this list (Serialization: [4,5,6])
```

## Solution

### Fast Pointer and Slow Pointer

The fast pointer moves two steps each time, while the slow pointer moves one step each time. When the fast pointer arrives at the end of the list, the slow pointer must be in the middle of the list.

#### Code

##### Go

```go
func MiddleNode(head *ListNode) *ListNode {
    middleNode := head

    for head != nil && head.Next != nil {
        head = head.Next.Next
        middleNode = middleNode.Next
    }

    return middleNode
}
```

#### Complexity Analysis

- **Time Complexity** : `O(n)`, where `n` is the number of nodes in the linked list.

- **Space Complexity** : `O(1)`. The extra space is only used by the slow pointer.
