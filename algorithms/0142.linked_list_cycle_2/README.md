# Linked List Cycle II

## Problem

Given the `head` of a linked list, return the node where the cycle begins. If there is no cycle, return `null`.
There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer.
Internally, `pos` is used to denote the index of the node that tail's `next` pointer is connected to (0-indexed). It is -1 if there is no cycle. Note that `pos` is not passed as a parameter.

Do not modify the linked list.

## Examples

```text
Input: head = [3,2,0,-4], pos = 1

Output: tail connects to node index 1
```

```text
Input: head = [1,2], pos = 0

Output: tail connects to node index 0
```

```text
Input: head = [1], pos = -1

Output: no cycle
```

## Solution

### Hash Table Solution

The hash table solution is straightforward. We traverse the linked list and store each node's address in a hash table. If the current node's address is already in the hash table, then we have a cycle. The last node in the cycle is the node whose address is the same as the current node's address.

#### Code

##### Go

```go
func DetectCycle(head *ListNode) *ListNode {
    nodesMap := make(map[*ListNode]*ListNode)

    currentNode := head

    for currentNode != nil {
        if node, ok := nodesMap[currentNode]; ok {
            return node
        }

        nodesMap[currentNode] = currentNode
        currentNode = currentNode.Next
    }

    return nil
}
```

#### Complexity Analysis

- **Time Complexity:** `O(n)`, where `n` is the number of nodes in the linked list. We visit each of the `n` nodes at most once by checking if the node has been visited before.

- **Space Complexity:** `O(n)`, the space used by the hash table.
