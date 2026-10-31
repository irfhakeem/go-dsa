package leetcode

import (
	"cmp"
	"slices"
)

// Time Complexity: O(n^2 log n)
// Space Complexity: O(n)
func sortMatrix(grid [][]int) [][]int {
	n := len(grid)

	start := 0
	for start < n {
		var tmpL []int

		i := start
		pointL := 0

		for i < n {
			tmpL = append(tmpL, grid[i][pointL])

			i++
			pointL++
		}

		slices.SortFunc(tmpL, func(a, b int) int {
			return cmp.Compare(b, a)
		})

		// fmt.Println(tmpL)

		i = start
		pointL = 0
		iL := 0

		for i < n {
			grid[i][pointL] = tmpL[iL]

			i++
			pointL++
			iL++
		}

		start++
	}

	start = n - 2
	for start >= 0 {
		var tmpR []int

		i := start
		pointR := n - 1

		for i >= 0 {
			tmpR = append(tmpR, grid[i][pointR])

			i--
			pointR--
		}

		slices.SortFunc(tmpR, func(a, b int) int {
			return cmp.Compare(a, b)
		})

		// fmt.Println(tmpR)

		i = start
		pointR = n - 1
		iR := len(tmpR) - 1

		for i >= 0 {
			grid[i][pointR] = tmpR[iR]

			i--
			pointR--
			iR--
		}

		start--
	}

	return grid
}
