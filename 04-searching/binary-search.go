package searching

import "slices"

func BinarySearch(arr []int, target int) int {
	if !slices.IsSorted(arr) {
		slices.Sort(arr)
	}

	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
