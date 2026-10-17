package leetcode

// Time Complexity: O(n)
// Space Complexity: O(1)
func splitArray(nums []int, k int) int {
	low := 0
	high := 0

	for _, num := range nums {
		if num > low {
			low = num
		}
		high += num
	}

	for low < high {
		mid := low + (high-low)/2
		totalSum := 0
		totaNum := 1

		for _, w := range nums {
			if totalSum+w > mid {
				totalSum = 0
				totaNum++
			}

			totalSum += w
		}

		if totaNum <= k {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}
