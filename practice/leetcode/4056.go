package leetcode

import (
	"cmp"
	"slices"
)

// Time Complexity: O(n^2)
// Space Complexity: O(1)
func countIntersectingIntervals(intervals [][]int) int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	count := 0
	l := 0
	r := l + 1

	for l < len(intervals)-1 {
		for r < len(intervals) {
			if intervals[l][1] >= intervals[r][0] {
				count++
				r++
			} else {
				break
			}
		}

		l++
		r = l + 1
	}

	return count
}
