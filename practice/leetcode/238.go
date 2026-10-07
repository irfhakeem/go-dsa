package leetcode

// Time Complexity: O(n)
// Space Complexity: O(n)
func productExceptSelf(nums []int) []int {
	n := len(nums)

	if n < 3 {
		return []int{nums[1], nums[0]}
	}

	var ans []int
	p := make([]int, n)
	s := make([]int, n)

	p[0] = 1
	s[n-1] = 1

	l := 0
	for l < n-1 {
		p[l+1] = p[l] * nums[l]
		l++
	}

	r := n - 1
	for r > 0 {
		s[r-1] = s[r] * nums[r]
		r--
	}

	for i := 0; i < n; i++ {
		ans = append(ans, p[i]*s[i])
	}

	return ans
}
