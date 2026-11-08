package leetcode

import (
	"cmp"
	"slices"
)

// Time Complexity: O(n log n)
// Space Complexity: O(n)
func maxKDistinct(nums []int, k int) []int {
	slices.SortFunc(nums, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	var ans []int
	choosed := make(map[int]bool)
	for _, num := range nums {
		_, ok := choosed[num]

		if !ok && k > 0 {
			choosed[num] = true
			ans = append(ans, num)
			k--
		}
	}

	return ans
}
