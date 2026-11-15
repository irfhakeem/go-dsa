package leetcode

// Time Complexity: O(n)
// Space Complexity: O(1)
func minMoves(nums []int) int {
	mx := -1
	for _, num := range nums {
		mx = max(mx, num)
	}

	ans := 0
	for _, num := range nums {
		ans += mx - num
	}

	return ans
}
