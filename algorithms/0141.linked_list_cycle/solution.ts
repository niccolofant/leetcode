interface ListNode {
  val: number
  next: ListNode | null
}

// 141. Linked List Cycle (Easy)
// https://leetcode.com/problems/linked-list-cycle

// Double pointer solution
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
