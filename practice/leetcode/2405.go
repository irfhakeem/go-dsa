package leetcode

import "fmt"

// Time Complexity: O(n)
// Space Complexity: O(n)
func partitionString(s string) int {
	seen := make(map[byte]bool)
	count := 1
	i := 0
	increment := 1

	for i < len(s) {
		seen[s[i]] = true
		for j := i + 1; j < len(s); j++ {
			fmt.Println(string(s[i]), string(s[j]))
			if _, ok := seen[s[j]]; ok {
				count++
				break
			}

			seen[s[j]] = true
			increment++
		}
		clear(seen)
		i += increment
		increment = 1
	}

	return count
}
