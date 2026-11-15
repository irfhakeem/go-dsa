package leetcode

func subarraySum(nums []int) int {
	n := len(nums)
	p := make([]int, n)
	p[0] = nums[0]

	for i := 1; i < n; i++ {
		p[i] = p[i-1] + nums[i]
	}

	ans := p[0]
	for i := 1; i < n; i++ {
		start := max(0, i-nums[i])

		if start > 0 {
			ans += p[i] - p[start-1]
		} else {
			ans += p[i]
		}
	}

	return ans
}
