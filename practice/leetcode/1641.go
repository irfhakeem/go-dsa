package leetcode

// Time Complexity: O(n)
// Space Complexity: O(n)
func countVowelStrings(n int) int {
	return ((n + 1) * (n + 2) * (n + 3) * (n + 4)) / 24
}
