# Missing Number

## Problem

Given an array `nums` containing `n` distinct numbers in the range `[0, n]`, return the only number in the range that is missing from the array.

## Example

```text
Input: nums = [3,0,1]

Output: 2
```

```text
Input: nums = [0,1]

Output: 2
```

## Solution

### Difference between total and sum of array

The total sum of the array is `n * (n + 1) / 2`. The sum of the array is `nums[0] + nums[1] + ... + nums[n - 1]`.
The difference between the total sum and the sum of the array is the missing number.

#### Code

##### Go

```go
func MissingNumber(nums []int) int {
    n := len(nums)
    expectedSumOfNums := (n * (n + 1)) / 2
    actualSumOfNums := 0

    for _, num := range nums {
        actualSumOfNums += num
    }

    return expectedSumOfNums - actualSumOfNums
}
```

#### Complexity Analysis

- **Time Complexity** : `O(n)`, where `n` is the number of elements in the array.

- **Space Complexity** : `O(1)`. The extra space is only used by the variables `n`, `expectedSumOfNums` and `actualSumOfNums`.
