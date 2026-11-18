package leetcode

// Time Complexity: O(n^2)
// Space Complexity: O(1)
func hasSameDigits(s string) bool {
	var nums []int
	for i := 0; i < len(s); i++ {
		nums = append(nums, int(s[i]))
	}

	end := len(s) - 1
	for len(nums) > 2 {
		for i := 0; i < end; i++ {
			nums[i] = (nums[i] + nums[i+1]) % 10
		}
		nums = nums[:end]
		end--
	}
	a
	return nums[0] == nums[1]
}
