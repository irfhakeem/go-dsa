package arraysandstrings

import (
	"regexp"
	"strings"
)

// Time Complexity: O(n)
func IsPalindrome(s string) bool {
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	cleaned := re.ReplaceAllString(s, "")
	cleaned = strings.ToLower(cleaned)

	i := 0
	j := len(cleaned) - 1
	for i < j {
		if cleaned[i] != cleaned[j] {
			return false
		}
		i++
		j--
	}
	return true
}
