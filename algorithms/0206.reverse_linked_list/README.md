# Reverse Linked List

## Problem

Reverse a singly linked list.
Given `head` of a singly linked list, reverse it and return the new head.

## Example

```text
Input: 1->2->3->4->5->NULL

Output: 5->4->3->2->1->NULL
```

```text
Input: NULL

Output: NULL
```

## Solution

### Iterative Approach

The iterative approach uses three pointers, `prev`, `curr`, and `next`, to reverse the linked list.
The `prev` pointer points to the previous node of the current node.
The `curr` pointer points to the current node.
The `next` pointer points to the next node of the current node.
The `prev` pointer is initialized to `nil` and the `curr` pointer is initialized to the head of the linked list.
The `next` pointer is initialized to the next node of the current node.
The `curr` pointer is then set to the `prev` pointer and the `prev` pointer is set to the `curr` pointer.
The `curr` pointer is then set to the `next` pointer and the `next` pointer is set to the next node of the current node.
The process is repeated until the `curr` pointer reaches the end of the linked list.
The `prev` pointer is then returned as the new head of the linked list.

#### Code

##### Go

```go
func reverseList(head *ListNode) *ListNode {
  var prev *ListNode = nil
  curr := head

  for curr != nil {
    next := curr.Next
    curr.Next = prev
    prev = curr
    curr = next
 }

  return prev
}
```

#### Complexity Analysis

- **Time Complexity** : `O(n)`, where `n` is the number of nodes in the linked list.

- **Space Complexity** : `O(1)`, since we only use three pointers.
