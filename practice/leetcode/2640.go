package leetcode

// Time Complexity: O(n)
// Space Complexity: O(n)
func findPrefixScore(nums []int) []int64 {
	ans := make([]int64, len(nums))
	mx := nums[0]
	ans[0] = int64(nums[0] + mx)

	for i := 1; i < len(nums); i++ {
		mx = max(mx, nums[i])
		ans[i] = ans[i-1] + int64(nums[i]) + int64(mx)
	}

	return ans
}
