package leetcode

// Time Complexity: O(n)
// Space Complexity: O(n)
func maxDistinct(s string) int {
	used := make(map[rune]bool)
	runes := []rune(s)
	count := 0

	for _, c := range runes {
		_, ok := used[c]

		if !ok {
			used[c] = true
			count++
		}
	}

	return count
}
