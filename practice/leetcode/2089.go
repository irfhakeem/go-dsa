package leetcode

// Time Complexity: O(n)
// Space Complexity: (n)
func targetIndices(nums []int, target int) []int {
	targetCount := 0
	lessCount := 0
	for _, num := range nums {
		if num < target {
			lessCount++
		} else if num == target {
			targetCount++
		}
	}

	ans := make([]int, targetCount)
	for i := lessCount; i < targetCount+lessCount; i++ {
		ans[i-lessCount] = i
	}

	return ans
}
