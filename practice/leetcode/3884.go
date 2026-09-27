package leetcode

// Time Complexity: O(n)
// Space Complexity: O(1)
func firstMatchingIndex(s string) int {
	n := len(s)

	for i, _ := range s {
		if s[i] == s[n-i-1] {
			return i
		}
	}

	return -1
}
