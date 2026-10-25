package leetcode

// Time Complexity: O(n)
// Space Complexity: O(n)
func countPartitions(nums []int) int {
	n := len(nums)
	preff := make([]int, n)
	preff[0] = nums[0]

	for i := 1; i < n; i++ {
		preff[i] = preff[i-1] + nums[i]
	}

	// fmt.Println(preff)

	count := 0
	for i := 0; i < n-1; i++ {
		if (preff[i]-(preff[n-1]-preff[i]))%2 == 0 {
			count++
		}
	}

	return count
}
