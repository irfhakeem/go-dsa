package leetcode

import "slices"

// Time Complexity: O(n log n)
// Space Complexity: O(n)
func reductionOperations(nums []int) int {
	slices.Sort(nums)

	count := 0
	add := 0
	i := 1

	for i < len(nums) {
		if nums[i] != nums[i-1] {
			add++
		}

		count += add
		i++
	}

	return count
}
