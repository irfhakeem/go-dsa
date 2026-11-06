package leetcode

import (
	"cmp"
	"slices"
)

// Time Complexity: O(n log n)
// Space Complexity: O(1)
func absDifference(nums []int, k int) int {
	slices.SortFunc(nums, func(a, b int) int {
		return cmp.Compare(a, b)
	})

	countL := 0
	countR := 0
	l := 0
	r := len(nums) - 1

	for _ = range k {
		countL += nums[l]
		countR += nums[r]
		l++
		r--
	}

	if countL > countR {
		return countL - countR
	}

	return countR - countL
}
