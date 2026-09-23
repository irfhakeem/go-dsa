package leetcode

// Time Complexity: O(n)
// Space Complexity: O(1)
func getDescentPeriods(prices []int) int64 {
	if len(prices) < 2 {
		return 1
	}

	total := int64(0)
	count := int64(1)

	for i := range len(prices) - 1 {
		if prices[i]-prices[i+1] == 1 {
			count++
		} else if prices[i]-prices[i+1] != 1 && count > 1 {
			total += (count * (count + int64(1))) / int64(2)
			count = int64(1)
		} else {
			total++
		}
	}

	if count > 1 {
		total += (count * (count + int64(1))) / int64(2)
	} else {
		total++
	}

	return total
}
