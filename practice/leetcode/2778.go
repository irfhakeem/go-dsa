package leetcode

// Time Complexity: O(n)
// Space Complexity: O(1)
func sumOfSquares(nums []int) int {
	n := len(nums)
	ans := 0

	for i, num := range nums {
		if n%(i+1) == 0 {
			ans += (num * num)
		}
	}

	return ans
}
