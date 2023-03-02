package algorithms

// 9. Palindrome Number (Easy)
// https://leetcode.com/problems/palindrome-number/

// Modulo and division solution
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