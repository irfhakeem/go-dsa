package leetcode

import (
	"cmp"
	"slices"
)

// Time Complexity: O(n log n)
// Space Complexity: O(n)
func divideArray(nums []int, k int) [][]int {
	slices.SortFunc(nums, func(a, b int) int {
		return cmp.Compare(a, b)
	})

	var ans [][]int
	for i := 0; i < len(nums); i += 3 {
		if nums[i+2]-nums[i] <= k {
			ans = append(ans, nums[i:i+3])
		}
	}

	if len(ans) != len(nums)/3 {
		return nil
	}

	return ans
}
