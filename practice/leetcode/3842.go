package leetcode

import "slices"

// Time Complexity: O(n log n)
// Space Complexity: O(n)
func toggleLightBulbs(bulbs []int) []int {
	store := make(map[int]bool)
	for _, b := range bulbs {
		store[b] = store[b] == false
	}

	var ans []int
	for k, v := range store {
		if v {
			ans = append(ans, k)
		}
	}

	slices.Sort(ans)

	return ans
}
