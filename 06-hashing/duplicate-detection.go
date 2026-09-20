package hashing

func DuplicateDetection(nums []int) int {
	// Time Complexity: O(n)
	// Space Complexity: O(n)
	seen := make(map[int]int)
	for _, num := range nums {
		seen[num]++
		if seen[num] > 1 {
			return num
		}
	}
	return -1

	// // Time Complexity: O(n^2)
	// // Space Complexity: O(1)
	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i] == nums[j] {
	// 			return nums[i]
	// 		}
	// 	}
	// }
	// return -1

	// // Time Complexity: O(n log n)
	// // Space Complexity: O(1)
	// slices.Sort(nums)
	// for i := 0; i < len(nums)-1; i++ {
	// 	if nums[i] == nums[i+1] {
	// 		return nums[i]
	// 	}
	// }
	// return -1
}
