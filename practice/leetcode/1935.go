package leetcode

import "strings"

// Time Complexity: O(n + m)
// Space Complexity: O(n + m)
func canBeTypedWords(text string, brokenLetters string) int {
	words := strings.Split(text, " ")

	letters := make([]int, 26)
	for i, _ := range brokenLetters {
		letters[brokenLetters[i]-'a'] = 1
	}

	count := len(words)
	for _, w := range words {
		for _, c := range w {
			if letters[c-'a'] == 1 {
				count--
				break
			}
		}
	}

	return count
}
