package leetcode

import "slices"

// Time Complexity: O(n log n)
// Space Complexity: O(1)
func countPairs(nums []int, target int) int {
	slices.Sort(nums)

	l := 0
	r := len(nums) - 1
	count := 0

	for l < r {
		if nums[l]+nums[r] < target {
			count += r - l
			l++
		} else {
			r--
		}
	}

	return count
}
