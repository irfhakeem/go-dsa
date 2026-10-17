package leetcode

// Time Complexity: O(n)
// Space Complexity: O(1)
func doesAliceWin(s string) bool {
	count := 0

	for _, ch := range s {
		if ch == 'a' || ch == 'i' || ch == 'u' || ch == 'e' || ch == 'o' {
			count++
		}
	}

	if count == 0 {
		return false
	}

	return true
}
