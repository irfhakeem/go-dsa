package leetcode

import "slices"

// Time Complexity: O(n)
// Space Complexity: O(1)
func limitOccurrences(nums []int, k int) []int {
	cn := nums[0]
	tv := 1
	i := 1

	for i < len(nums) {
		if cn == nums[i] && tv < k {
			tv++
			i++
		} else if cn == nums[i] && tv >= k {
			nums = slices.Delete(nums, i, i+1)
		} else {
			cn = nums[i]
			tv = 1
			i++
		}
	}

	return nums
}
