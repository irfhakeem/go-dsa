package leetcode

import "strings"

// Time Complexity: O(n)
// Space Complexity: O(1)
func countKeyChanges(s string) int {
	count := 0
	i := 0
	j := i + 1

	for i < len(s)-1 {
		if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
			count++
		}

		i++
		j++
	}

	return count
}
