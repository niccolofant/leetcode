# Binary Search

## Problem

Given an array of integers `nums` which is sorted in ascending order, and an integer `target`, write a function to search `target` in `nums`. If `target` exists, then return its index. Otherwise, return `-1`.
The algorithm must be in the order of `O(log n)`.

## Examples

```text
Input: nums = [-1,0,3,5,9,12], target = 9

Output: 4
```

```text
Input: nums = [-1,0,3,5,9,12], target = 2

Output: -1
```

## Solution

### Recursive Solution

We keep track of the left (`low`) and right (`high`) indices of the array.
If the `low` index is greater than the `high` index, we return `-1`.
Otherwise, we calculate the middle index (`mid`) of the array.
If the value at the `mid` index is equal to `target`, we return the `mid` index.
Otherwise, if the value at the `mid` index is greater than `target`, we search the left half of the array.
Otherwise, we search the right half of the array.

#### Code

##### Go

```go
func Search(nums []int, target int) int {
    return BinarySearch(nums, 0, len(nums)-1, target)
}

func BinarySearch(nums []int, low int, high int, target int) int {
    if low <= high {
        mid := (low + high) / 2

        if target == nums[mid] {
            return mid
        } else {
            if target > nums[mid] {
                return BinarySearch(nums, mid+1, high, target)
            } else {
                return BinarySearch(nums, low, mid-1, target)
            }
        }
    }

    return -1
}
```

#### Complexity Analysis

- **Time Complexity** : `O(log n)`, where `n` is the number of elements in the array.

- **Space Complexity** : `O(log n)`, where `n` is the number of elements in the array.

### Iterative Solution

We keep track of the left (`low`) and right (`high`) indices of the array.
We loop while the `low` index is less than or equal to the `high` index.
We calculate the middle index (`mid`) of the array.
If the value at the `mid` index is equal to `target`, we return the `mid` index.
Otherwise, if the value at the `mid` index is greater than `target`, we search the left half of the array.
Otherwise, we search the right half of the array.

#### Code

##### Go

```go
func Search(nums []int, target int) int {
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

    return -1
}
```

#### Complexity Analysis

- **Time Complexity** : `O(log n)`, where `n` is the number of elements in the array.

- **Space Complexity** : `O(1)`, since we are not using any extra space.
