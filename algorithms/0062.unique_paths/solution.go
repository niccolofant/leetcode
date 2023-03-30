package algorithms

import "fmt"

// 62. Unique Paths (Medium)
// https://leetcode.com/problems/unique-paths

// Dynamic Programming - Memoization solution
func UniquePaths(m int, n int) int {
	return UniquePathsWithMemo(m, n, 0, 0, make(map[string]int))
}

func UniquePathsWithMemo(m int, n int, i int, j int, seen map[string]int) int {
	mapKey := computeMapKey(i, j)

	if v, ok := seen[mapKey]; ok {
		return v
	}

	if i == m-1 && j == n-1 {
		seen[mapKey] = 1

		return seen[mapKey]
	}

	if i == m-1 {
		seen[mapKey] = UniquePathsWithMemo(m, n, i, j+1, seen)

		return seen[mapKey]
	}

	if j == n-1 {
		seen[mapKey] = UniquePathsWithMemo(m, n, i+1, j, seen)

		return seen[mapKey]
	}

	seen[mapKey] = UniquePathsWithMemo(m, n, i, j+1, seen) + UniquePathsWithMemo(m, n, i+1, j, seen)

	return seen[mapKey]
}

// Example: 10, 5 -> "10:5"
func computeMapKey(a, b int) string {
	return fmt.Sprintf("%d:%d", a, b)
}
