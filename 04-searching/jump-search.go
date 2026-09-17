package searching

import "slices"

// Time Complexity: O(n log n)
func JumpSearch(arr []int, target int, jumpSize int) int {
	slices.Sort(arr)

	i := 0
	var lowerBound int
	var upperBound int

	// search for target range
	for i < len(arr) {
		if arr[i] < target {
			lowerBound = i
		}

		if arr[i] > target {
			upperBound = i
			break
		}

		if arr[i] == target {
			return arr[i]
		}

		i += jumpSize
	}

	// search within target range
	for j := lowerBound; j < upperBound; j++ {
		if arr[j] == target {
			return arr[j]
		}
	}

	return -1
}
