# Find the Pivot Integer

## Problem

Given a positive integer `n`, find the pivot integer `k` such that:

- The sum of all elements between `1` and `k` inclusively equals the sum of all elements between `k` and `n` inclusively.

Return the pivot integer `k`. If no such integer exists, return `-1`. It is guaranteed that there will be at most one pivot index for the given input.

## Example

```text
Input: n = 4

Output: 2
```

```text
Input: n = 6

Output: 3
```

```text
Input: n = 2

Output: -1
```

## Solution

### Sum and Subtraction Approach

We keep track of the sum of all elements between `1` and `n` inclusively.
We also keep track of the sum of all elements between `1` and `k` inclusively.
We then loop through all integers between `1` and `n` inclusively.
For each integer, we subtract the current integer from the sum of all elements between `1` and `n` inclusively.
We then add the current integer to the sum of all elements between `1` and `k` inclusively.
If the sum of all elements between `1` and `k` inclusively equals the sum of all elements between `k` and `n` inclusively, we return the current integer.
If we reach the end of the loop, we return `-1`.

#### Code

##### Go

```go
func pivotInteger(n int) int {
  sumOfAllNumbers := n * (n + 1) / 2
  sumOfFirstKNumbers := 0

  k := 1

  for k <= n {
    sumOfFirstKNumbers += k

    if sumOfFirstKNumbers == sumOfAllNumbers {
      return k
    }

    sumOfAllNumbers -= k

    k++
  }

  return -1
}
```

##### TypeScript

```typescript
const pivotIndex = (n: number): number => {
  let sumOfAllNumbers = (n * (n + 1)) / 2
  let sumOfFirstKNumbers = 0

  for (let k = 1; k <= n; k++) {
    sumOfFirstKNumbers += k

    if (sumOfFirstKNumbers === sumOfAllNumbers) {
      return k
    }

    sumOfAllNumbers -= k
  }

  return -1
}
```

#### Complexity Analysis

- **Time Complexity** : `O(n)` where `n` is the input integer.

- **Space Complexity** : `O(1)`.
