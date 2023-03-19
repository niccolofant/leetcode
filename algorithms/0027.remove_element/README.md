# Remove Element

## Problem

Given an array `nums` and a value `val`, remove all instances of that value in-place and return the new length.
The order of elements can be changed. It doesn't matter what you leave beyond the new length.
Do not allocate extra space for another array. You must do this by modifying the input array in-place with `O(1)` extra memory.

## Examples

```text
Input: nums = [3,2,2,3], val = 3

Output: 2
```

```text
Input: nums = [0,1,2,2,3,0,4,2], val = 2

Output: 5
```

## Solution

### Two Pointers Solution

We initialize two pointer `i` and `k` to `0` and `len(nums)-1` respectively.
We iterate through the array while `i` is less than or equal to `k`.
If the value at index `i` is equal to `val`, we swap the value at index `i` with the value at index `k` and decrement `k`.
Otherwise, we increment `i`.
At the end of the iteration, we return `k+1`.

#### Code

##### Go

```go
func RemoveElement(nums []int, val int) int {
    k := len(nums) - 1
    i := 0

    for i <= k {
        if nums[i] == val {
            nums[i] = -1 // Not really necessary
            nums[i], nums[k] = nums[k], nums[i]
            k--
        } else {
            i++
        }
    }

    return k + 1
}
```

#### Complexity Analysis

- **Time Complexity** : `O(n)` where `n` is the number of elements in the array. Because we are iterating through the array once.

- **Space Complexity** : `O(1)`. Because we are not using any extra space.
