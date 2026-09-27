package leetcode

import (
	"math"
	"slices"
)

// Time Complexity: O(n log n + m log n)
// Space Complexity: O(1)
func minPrice(prices []int, discounts []int) float64 {
	slices.Sort(discounts)
	slices.SortFunc(prices, func(a, b int) int {
		return b - a
	})

	var ans float64
	i := len(discounts) - 1
	j := 0

	for i >= 0 {
		if j < len(prices) {
			ans += math.Round(
				float64(prices[j])*float64(100-discounts[i])/100*1e5,
			) / 1e5
		}

		i--
		j++
	}

	for j < len(prices) {
		ans += float64(prices[j])
		j++
	}

	return math.Round(ans*1e5) / 1e5
}
