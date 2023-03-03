# Linked List Cycle

## Problem

Given `head`, the head of a linked list, determine if the linked list has a cycle in it.
There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer.
Internally, `pos` is used to denote the index of the node that tail's `next` pointer is connected to.
**Note that** `pos` **is not passed as a parameter.**

Return `true` if there is a cycle in the linked list. Otherwise, return `false`.

## Example

```text
Input: head = [3,2,0,-4], pos = 1

Output: true
```

```text
Input: head = [1,2], pos = 0

Output: true
```

```text
Input: head = [1], pos = -1

Output: false
```

## Solution

### Double Pointer Approach

The Double Pointer approach uses two pointers, `slow` and `fast`, to determine if there is a cycle in the linked list.
The `slow` pointer moves one step at a time while the `fast` pointer moves two steps at a time.
If there is no cycle in the linked list, the `fast` pointer will eventually reach the end and we can return `false` in this case.
Otherwise, the `fast` pointer will eventually meet the `slow` pointer and we can return `true` in this case.

#### Code

##### Go

```go
func hasCycle(head *ListNode) bool {
  if head == nil || head.Next == nil {
    return false
  }

  slowHead := head
  fastHead := head

  for fastHead != nil && fastHead.Next != nil {
    slowHead = slowHead.Next
    fastHead = fastHead.Next.Next

    if fastHead == slowHead {
      return true
    }
  }

  return false
}
```

##### TypeScript

```typescript
const hasCycle = (head: ListNode | null): boolean => {
  if (!head) return false

  let slowHead: ListNode | null | undefined = head
  let fastHead: ListNode | null | undefined = head

  while (slowHead && fastHead) {
    slowHead = slowHead.next
    fastHead = fastHead.next?.next

    if (slowHead === fastHead) return true
  }

  return false
}
```

#### Complexity Analysis

- **Time Complexity** : `O(n)`.

  - The algorithm has a linear time complexity because the `fast` pointer is moving at twice the speed as the `slow` pointer and the distance between the `slow` and `fast` pointers is cut in half at every iteration of the loop.

- **Space Complexity** : `O(1)`.
