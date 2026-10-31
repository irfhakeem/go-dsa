package leetcode

// Time Complexity: O(n)
// Space Complexity: O(n)
func recoverOrder(orders []int, friends []int) []int {
	stores := make(map[int]bool)

	var ans []int
	for _, f := range friends {
		stores[f] = true
	}

	for _, o := range orders {
		if _, ok := stores[o]; ok {
			ans = append(ans, o)
		}
	}

	return ans
}
