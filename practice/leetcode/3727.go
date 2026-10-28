package leetcode

import (
	"cmp"
	"math"
	"slices"
)

// Time Complexity: O(n)
// Space Complexity: O(1)
func maxAlternatingSum(nums []int) int64 {
	var store []int64
	for _, num := range nums {
		tmp := int64(num) * int64(num)
		store = append(store, tmp)
	}

	// fmt.Println(store)

	totalPlus := math.Ceil(float64(len(store)) / float64(2))

	slices.SortFunc(store, func(a, b int64) int {
		return cmp.Compare(b, a)
	})

	// fmt.Println(store)

	var count int64
	for _, num := range store {
		if totalPlus > 0 {
			count += num
			totalPlus--
			continue
		}

		count -= num
	}

	return count
}
