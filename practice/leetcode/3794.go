package leetcode

// Time Complexity: O(k)
// Space Complexity: O(n)
func reversePrefix(s string, k int) string {
	if k == 1 {
		return s
	}

	runes := []rune(s)

	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
