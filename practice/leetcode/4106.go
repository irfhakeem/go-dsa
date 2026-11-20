package leetcode

import "strings"

// Time Complexity: O(n)
// Space Complexity: O(n)
func findMismatchedSegments(s1 string, s2 string) [][]string {
	var ans [][]string
	i := 0
	for i < len(s1) {
		var b1, b2 strings.Builder

		if s1[i] != s2[i] {
			b1.WriteByte(s1[i])
			b2.WriteByte(s2[i])

			for j := i + 1; j < len(s1); j++ {
				i = j
				if s1[j] != s2[j] {
					b1.WriteByte(s1[j])
					b2.WriteByte(s2[j])
				} else {
					break
				}
			}
		}

		if b1.Len() > 0 {
			ans = append(ans, []string{
				b1.String(),
				b2.String(),
			})
		}

		i++
	}

	return ans
}
