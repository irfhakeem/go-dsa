package leetcode

// Time Complexity: O(nxm)
// Space Complexity: O(1)
func prefixCount(words []string, pref string) int {
	n := len(pref)
	count := 0
	for _, w := range words {
		if len(w) >= n {
			if w[:n] == pref {
				count++
			}
		}
	}

	return count
}
