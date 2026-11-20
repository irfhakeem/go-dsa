package leetcode

// Time Complexity: O(n)
// Space Complexity: O(1)
func minLengthAfterRemovals(s string) int {
	c := 0
	for i, _ := range s {
		if s[i] == 'a' {
			c++
		} else {
			c--
		}
	}

	return max(c, -c)
}
