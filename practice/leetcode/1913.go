package leetcode

import (
	"cmp"
	"slices"
)

// Time Complexity: O(n log n)
// Space Comlexity: O(1)
func maxProductDifference(nums []int) int {
	slices.SortFunc(nums, func(a, b int) int {
		return cmp.Compare(a, b)
	})

	// fmt.Println(nums)

	last := len(nums) - 1
	start := 0

	return (nums[last] * nums[last-1]) - (nums[start] * nums[start+1])
}
