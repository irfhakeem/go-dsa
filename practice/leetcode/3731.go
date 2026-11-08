package leetcode

// Time Complexity: O(n)
// Space Complexity: O(n)
func findMissingElements(nums []int) []int {
	mn := 101
	mx := -1
	mp := make(map[int]bool)
	for _, num := range nums {
		mn = min(mn, num)
		mx = max(mx, num)
		mp[num] = true
	}

	var ans []int
	for i := mn + 1; i < mx; i++ {
		if _, ok := mp[i]; !ok {
			ans = append(ans, i)
		}
	}

	return ans
}
