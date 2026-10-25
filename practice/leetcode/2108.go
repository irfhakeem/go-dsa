package leetcode

// Time Complexity: O(n^m)
// Space Complexity: O(1)
func firstPalindrome(words []string) string {
	for _, w := range words {
		isPalindromic := true
		l := 0
		r := len(w) - 1

		for l < r {
			if w[l] != w[r] {
				isPalindromic = false
				break
			}
			l++
			r--
		}

		if isPalindromic {
			return w
		}
	}

	return ""
}
