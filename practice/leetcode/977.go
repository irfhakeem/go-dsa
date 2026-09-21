package leetcode

import (
	"math"
	"slices"
)

func sortedSquares(nums []int) []int {
	for i, num := range nums {
		nums[i] = int(math.Pow(float64(num), float64(2)))
	}

	slices.Sort(nums)

	return nums
}
