package leetcode

import (
	"cmp"
	"slices"
)

// Time Complexity: O(r log r)
// Space Complexity: O(1)
func maxBuilding(n int, restrictions [][]int) int {
	restrictions = append(restrictions, []int{1, 0})

	slices.SortFunc(restrictions, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	if restrictions[len(restrictions)-1][0] != n {
		restrictions = append(restrictions, []int{n, n - 1})
	}

	for i := 1; i < len(restrictions); i++ {
		prev := restrictions[i-1]
		curr := restrictions[i]

		maxAllowed := prev[1] + curr[0] - prev[0]

		curr[1] = min(curr[1], maxAllowed)
	}
	// fmt.Println(restrictions)

	for i := len(restrictions) - 2; i >= 0; i-- {
		curr := restrictions[i]
		next := restrictions[i+1]

		maxAllowed := next[1] + next[0] - curr[0]

		curr[1] = min(curr[1], maxAllowed)
	}
	// fmt.Println(restrictions)

	mx := 0

	for i := 1; i < len(restrictions); i++ {
		prev := restrictions[i-1]
		curr := restrictions[i]

		distance := curr[0] - prev[0]

		peak := (prev[1] + curr[1] + distance) / 2

		mx = max(mx, peak)
	}

	return mx
}
