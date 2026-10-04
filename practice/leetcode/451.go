package leetcode

import (
	"sort"
	"strings"
)

// Time Complexity: O(n)
// Space Complexity: O(n)
func frequencySort(s string) string {
	chs := make(map[rune]int)

	for _, ch := range s {
		chs[ch]++
	}

	mp := make(map[int][]rune)

	for ch, total := range chs {
		mp[total] = append(mp[total], ch)
	}

	totals := make([]int, 0, len(mp))

	for total := range mp {
		totals = append(totals, total)
	}

	sort.Slice(totals, func(i, j int) bool {
		return totals[i] > totals[j]
	})

	ans := ""

	for _, total := range totals {
		for _, ch := range mp[total] {
			ans += strings.Repeat(string(ch), total)
		}
	}

	return ans
}
