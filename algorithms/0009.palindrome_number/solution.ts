// 9. Palindrome Number (Easy)
// https://leetcode.com/problems/palindrome-number/

// Modulo and division solution
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
