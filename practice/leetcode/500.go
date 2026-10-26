package leetcode

import "strings"

// Time Complexity: O(n^m)
// Space Complexity: O(n)
func findWords(words []string) []string {
	key := []int{1, 2, 2, 1, 0, 1, 1, 1, 0, 1, 1, 1, 2, 2, 0, 0, 0, 0, 1, 0, 0, 2, 0, 2, 0, 2}

	var ans []string
	for _, word := range words {
		lowerW := strings.ToLower(word)
		currRow := key[lowerW[0]-'a']
		isOneRow := true

		for _, c := range lowerW {
			if key[c-'a'] != currRow {
				isOneRow = false
				break
			}
		}

		if isOneRow {
			ans = append(ans, word)
		}
	}

	return ans
}
