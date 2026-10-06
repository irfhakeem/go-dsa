package hackerrank

import "slices"

// Time Complexity: O(n log n)
// Space Complexity: O(1)
func findMedian(arr []int32) int32 {
	// Write your code here
	slices.Sort(arr)

	mid := len(arr) / 2

	return arr[mid]
}
