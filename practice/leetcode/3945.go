package leetcode

import "strconv"

// Time Complexity: O(n)
// Space Complexity: O(1)
func digitFrequencyScore(n int) int {
	ans := 0
	sNum := strconv.Itoa(n)

	for _, c := range sNum {
		iNum, _ := strconv.Atoi(string(c))
		ans += iNum
	}

	return ans
}
