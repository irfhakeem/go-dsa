package leetcode

import "strings"

// Time Complexity: O(n + m)
// Space Complexity: O(m)
func customSortString(o string, s string) string {
	mp := make(map[rune]int)
	for _, c := range s {
		mp[c]++
	}

	var ans string
	for _, c := range o {
		if _, ok := mp[c]; ok {
			ans += strings.Repeat(string(c), mp[c])
			mp[c] = 0
		}
	}

	for k, v := range mp {
		if v != 0 {
			ans += strings.Repeat(string(k), v)
			mp[k] = 0
		}
	}

	return ans
}
