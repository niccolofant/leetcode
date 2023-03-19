# Search Insert Position

## Problem

Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
The algorithm must be in the order of `O(log n)`.

## Examples

```text
Input: nums = [1,3,5,6], target = 5

Output: 2
```

```text
Input: nums = [1,3,5,6], target = 2

Output: 1
```

```text
Input: nums = [1,3,5,6] target = 7

Output: 4
```

## Solution

### Binary Search Solution

We keep track of the left (`low`) and right (`high`) indices of the array.
If the `low` index is greater than the `high` index, we return the `low` index.
Otherwise, we calculate the middle index (`mid`) of the array.
If the value at the `mid` index is equal to `target`, we return the `mid` index.
Otherwise, if the value at the `mid` index is greater than `target`, we search the left half of the array.
Otherwise, we search the right half of the array.

#### Code

##### Go

```go
func SearchInsert(nums []int, target int) int {
    low := 0
    high := len(nums) - 1

    for low <= high {
        mid := (low + high) / 2

        if target == nums[mid] {
            return mid
        } else {
            if target > nums[mid] {
                low = mid + 1
            } else {
                high = mid - 1
            }
        }
    }

    return low
}
```

#### Complexity Analysis

- **Time Complexity** : `O(log n)` where `n` is the length of the array.

- **Space Complexity** : `O(1)`.
