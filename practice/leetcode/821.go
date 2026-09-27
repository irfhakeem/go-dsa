package leetcode

import "math"

// Time Complexity: O(n^2)
// Space Complexity: O(n)
func shortestToChar(s string, c byte) []int {
	var eIdx []int

	for i, ch := range s {
		if byte(ch) == c {
			eIdx = append(eIdx, i)
		}
	}

	var ans []int

	for i, _ := range s {
		mn := int(math.Abs(float64(eIdx[0]) - float64(i)))
		j := 1

		for j < len(eIdx) {
			mn = min(mn, int(math.Abs(float64(eIdx[j])-float64(i))))
			j++
		}

		ans = append(ans, mn)
	}

	return ans
}
