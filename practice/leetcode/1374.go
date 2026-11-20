package leetcode

import "strings"

// Time Complexity: O(n)
// Space Complexity: O(1)
func generateTheString(n int) string {
	var b strings.Builder
	for _ = range n - 1 {
		b.WriteByte('a')
	}

	if n%2 != 0 {
		b.WriteByte('a')
	} else {
		b.WriteByte('b')
	}

	return b.String()
}
