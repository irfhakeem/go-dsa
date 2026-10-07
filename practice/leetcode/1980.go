package leetcode

// Time Complexity: O(n)
// Space Compliexy: O(1)
func findDifferentBinaryString(nums []string) string {
	var ans string

	for i, _ := range nums {
		if nums[i][i] == '0' {
			ans += "1"
		} else {
			ans += "0"
		}
	}

	return ans
}
