package leetcode

// Time Complexity: O(n)
// Space Complexity: O(n)
func firstStableIndex(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	pref := make([]int, n)
	suff := make([]int, n)

	prev := nums[0]
	pref[0] = nums[0]
	for i := 1; i < n; i++ {
		if prev < nums[i] {
			prev = nums[i]
		}
		pref[i] = prev
	}

	prev = nums[n-1]
	suff[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if prev > nums[i] {
			prev = nums[i]
		}
		suff[i] = prev
	}

	// fmt.Println(pref, suff)

	for i := range n {
		tmp := pref[i] - suff[i]
		if tmp <= k {
			return i
		}
	}

	return -1
}
