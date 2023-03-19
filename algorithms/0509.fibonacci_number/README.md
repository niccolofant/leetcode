# Fibonacci Number

## Problem

The Fibonacci numbers, commonly denoted `F(n)` form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from `0` and `1`. That is,

```text
F(0) = 0, F(1) = 1

F(n) = F(n - 1) + F(n - 2), for n > 1
```

Given `n`, calculate `F(n)`.

## Examples

```text
Input: n = 2

Output: 1
```

```text
Input: n = 3

Output: 2
```

```text
Input: n = 4

Output: 3
```

## Solution

### Recursive Solution with Memoization

We use a map to store the values of the Fibonacci numbers that we have already calculated.
If the current value of `n` is found in the map, we return the value from the map.
If the current value of `n` is not found in the map, we calculate the value of `F(n)` recursively and store it in the map.

#### Code

##### Go

```go
func Fib(n int) int {
    return FibMemo(n, make(map[int]int))
}

func FibMemo(n int, memo map[int]int) int {
    if memoizedValue, ok := memo[n]; ok {
        return memoizedValue
    }

    if n == 0 {
        return 0
    }

    if n <= 2 {
        return 1
    }

    memo[n] = FibMemo(n-1, memo) + FibMemo(n-2, memo)

    return memo[n]
}
```

#### Complexity Analysis

- **Time Complexity** : `O(n)`, where `n` is the value of `n`.

- **Space Complexity** : `O(n)`. The extra space is used by the map to store the values of the Fibonacci numbers that we have already calculated.
