package leetcode

import (
	"cmp"
	"slices"
)

// Time Complexity: O(n log n)
// Space Complexity: O(1)
func maxIceCream(costs []int, coins int) int {
	slices.SortFunc(costs, func(a, b int) int {
		return cmp.Compare(a, b)
	})

	count := 0
	for _, c := range costs {
		if coins >= c {
			coins -= c
			count++
		} else {
			return count
		}
	}

	return count
}
