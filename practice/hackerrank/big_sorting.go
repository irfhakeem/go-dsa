package hackerrank

import "slices"

func bigSorting(unsorted []string) []string {
	// Write your code here
	slices.SortFunc(unsorted, func(a, b string) int {
		if len(a) != len(b) {
			return len(a) - len(b)
		}

		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})

	return unsorted
}
