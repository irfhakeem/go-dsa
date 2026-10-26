package leetcode

// Time Complexity: O(n)
// Space Complexity: O(1)
func missingNumber(nums []int) int {
	n := len(nums)
	targetSum := 0

	for i := range n {
		targetSum += i + 1
	}

	actualSum := 0

	for _, num := range nums {
		actualSum += num
	}

	return targetSum - actualSum
}
