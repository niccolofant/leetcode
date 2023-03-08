# Two Sum II - Input array is sorted

## Problem

Given a **1-indexed** array of integers `numbers` that is already **sorted in non-decreasing order**, find two numbers such that they add up to a specific `target` number. Let these two numbers be `numbers[index1]` and `numbers[index2]` where `1 <= index1 < index2 <= numbers.length`.

Return the indices of the two numbers, `index1` and `index2`, added by one as an integer array `[index1, index2]` of length 2.

The tests are generated such that there is exactly one solution. You may not use the same element twice.

Your solution must use only constant extra space.

## Example

```text
Input: numbers = [2,7,11,15], target = 9

Output: [1,2]
```

```text
Input: numbers = [2,3,4], target = 6

Output: [1,3]
```

```text
Input: numbers = [-1,0], target = -1

Output: [1,2]
```

## Solution

### Two Pointers Approach

The two pointers approach uses two pointers, `left` and `right`, to find the two numbers that add up to the `target`.
The `left` pointer points to the first number in the array and the `right` pointer points to the last number in the array.
The sum of the two numbers pointed to by the `left` and `right` pointers is then calculated.
If the sum is equal to the `target`, then the indices of the two numbers are returned (added by one).
If the sum is less than the `target`, then the `left` pointer is incremented to point to the next number.
If the sum is greater than the `target`, then the `right` pointer is decremented to point to the previous number.
The process is repeated until the two numbers are found.

#### Code

##### Go

```go
func TwoSum(numbers []int, target int) []int {
    var result []int

    left := 0
    right := len(numbers) - 1

    for left < right {
        sum := numbers[left] + numbers[right]

        if sum > target {
            right--
        } else {
            if sum < target {
                left++
            } else {
                result = []int{left + 1, right + 1}
                break
            }
        }
    }

    return result
}
```

#### Complexity Analysis

- **Time Complexity** : `O(n)`, where `n` is the length of the array.

- **Space Complexity** : `O(1)`, since we only use two pointers.
