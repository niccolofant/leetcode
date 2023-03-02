#  Palindrome Number

## Problem

Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

## Example

```text
Input: 121

Output: true
```

```text
Input: -121

Output: false
```

## Solution

### Modulo and Division Approach

The Modulo and Division approach we build up the reversed integer one digit at a time (We aren't checking for overflow here). Then, we compare the reversed integer with the original integer, if they are the same, then the number is a palindrome.
The reversing process is done by repeatedly "popping" the last digit off of `number` and "pushing" it to the back of the `reversed` integer. In the end, `number` will be 0 and the `reversed` integer will be the reverse of the original `number`.

#### Code

##### Go

```go
func isPalindrome(number int) bool {
  if number < 0 {
    return false
  }

  originalNumber := number
  reversedNumber := 0

  for number > 0 {
    remainder := number % 10
    reversedNumber = (reversedNumber * 10) + remainder
    number /= 10
  }

  return reversedNumber == originalNumber
}
```

##### TypeScript

```typescript
const isPalindrome = (num: number): boolean => {
  if (num < 0) return false

  const originalNum = num
  let reversedNum = 0

  while (num > 0) {
    const remainder = num % 10
    reversedNum = reversedNum * 10 + remainder
    num = Math.floor(num / 10)
  }

  return reversedNum === originalNum
}
```

#### Complexity Analysis

- **Time complexity** : `O(log10(n))`. We divided the input by 10 for every iteration, so the time complexity is `O(log10(n))`.
- **Space complexity** : `O(1)`.
