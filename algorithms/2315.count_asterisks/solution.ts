// 2315. Count Asterisks (Easy)
// https://leetcode.com/problems/count-asterisks/description/

// Iterative solution
const countAsterisks = (s: string): number => {
  const chars = [...s]
  let numberOfAsterisks = 0
  let currentPair = 0

  chars.forEach((char) => {
    if (char === "|") currentPair++
    if (currentPair % 2 !== 0) return
    if (char === "*") numberOfAsterisks++
  })

  return numberOfAsterisks
}
