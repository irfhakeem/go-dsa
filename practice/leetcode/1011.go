package leetcode

// Time Complexity: O(n)
// Space Complexity: O(1)
func shipWithinDays(weights []int, days int) int {
	low := 0
	high := 0

	for _, w := range weights {
		if w > low {
			low = w
		}
		high += w
	}

	for low < high {
		mid := low + (high-low)/2
		totalWeight := 0
		totalDay := 1

		for _, w := range weights {
			if totalWeight+w > mid {
				totalWeight = 0
				totalDay++
			}

			totalWeight += w
		}

		if totalDay <= days {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}
