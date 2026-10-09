package leetcode

// Time Complexity: O(n)
// Space Complexity: O(1)
func minOperationsMakeBinaryElementsEqual(nums []int) int {
	count := 0
	n := len(nums)
	i := 0

	for i < n {
		// fmt.Println(nums, i)
		if nums[i] == 0 && i < n-2 {
			nums[i] ^= 1
			nums[i+1] ^= 1
			nums[i+2] ^= 1
			count++
		} else if nums[i] == 0 && i >= n-2 {
			return -1
		}

		i++
	}

	return count
}
