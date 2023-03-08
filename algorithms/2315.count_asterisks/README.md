# Count Asterisks

## Problem

Given a string `s`, where every **two** consecutive vertical bars `'|'` are grouped
into a **pair**. In other words, the 1st and 2nd `'|'` make a pair, the 3rd and 4th `'|'` make a pair, and so forth.

Return the number of `'*'` in `s`, excluding the `'*'` between each pair of `'|'`.

Note that each `'|'` will belong to exactly one pair.

## Examples

```text
Input: s = "l|*e*et|c**o|*de|"
Output: 2
Explanation: The considered characters are underlined: "l|*e*et|c**o|*de|".
The characters between the first and second '|' are excluded from the answer.
Also, the characters between the third and fourth '|' are excluded from the answer.
There are 2 asterisks considered. Therefore, we return 2.
```

```text
Input: s = "iamprogrammer"
Output: 0
Explanation: In this example, there are no asterisks in s. Therefore, we return 0.
```

```text
Input: s = "yo|uar|e**|b|e***au|tifu|l"
Output: 5
Explanation: The considered characters are underlined: "yo|uar|e**|b|e***au|tifu|l". There are 5 asterisks considered. Therefore, we return 5.
```

## Solution

### Iterative Approach

The iterative approach uses a `numberOfAsterisks` variable to keep track of the number of asterisks in the string.
We also use a `currentPair` variable to keep track of whether we are currently in a pair of vertical bars.
We increase the `numberOfAsterisks` variable by 1 if we encounter an asterisk and we are not in a pair of vertical bars.

#### Code

##### Go

```go
func countAsterisks(s string) int {
  numberOfAsterisks := 0
  currentPair := 0

  for _, char := range s {
    if char == '|' {
      currentPair++
    }

    if currentPair % 2 != 0 {
      continue
    }

    if char == '*' {
      numberOfAsterisks++
    }
  }

  return numberOfAsterisks
}
```

#### Complexity Analysis

- **Time Complexity** : `O(n)`, where `n` is the length of the string `s`.

- **Space Complexity** : `O(1)`, since we only use two variables.
