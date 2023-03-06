package algorithms

// 2315. Count Asterisks (Easy)
// https://leetcode.com/problems/count-asterisks/description/

// Iterative solution
func countAsterisks(s string) int {
	numberOfAsterisks := 0
	currentPair := 0

	for _, char := range s {
		if char == '|' {
			currentPair++
		}

		if currentPair%2 != 0 {
			continue
		}

		if char == '*' {
			numberOfAsterisks++
		}
	}

	return numberOfAsterisks
}
