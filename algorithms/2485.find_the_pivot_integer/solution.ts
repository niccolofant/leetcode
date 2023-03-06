// 2485. Find the Pivot Integer (Easy)
// https://leetcode.com/problems/find-the-pivot-integer/

// Add and subtract solution
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
