package leetcode

import (
	"cmp"
	"slices"
)

// Time Complexity: O(n)
// Space Complexity: O(1)
func maximizeExpressionOfThree(nums []int) int {
	slices.SortFunc(nums, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	return nums[0] + nums[1] - nums[len(nums)-1]
}
