package searching

import "slices"

// Time Complexity: O(n log n)
func TernarySearch(arr []int, target int) int {
	slices.Sort(arr)

	lowerBound := 0
	upperBound := len(arr) - 1

	for lowerBound <= upperBound {
		mid1 := lowerBound + (upperBound-lowerBound)/3
		mid2 := upperBound - (upperBound-lowerBound)/3

		if arr[mid1] == target {
			return arr[mid1]
		}
		if arr[mid2] == target {
			return arr[mid2]
		}

		// check target value near the m1 or m2
		// and adjust the bounds accordingly
		if target < arr[mid1] {
			upperBound = mid1 - 1

		} else if target > arr[mid2] {
			lowerBound = mid2 + 1
		} else {
			lowerBound = mid1 + 1
			upperBound = mid2 - 1
		}
	}

	return -1
}
