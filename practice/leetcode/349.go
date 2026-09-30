package leetcode

// Time Complexity: O(n+m)
// Space Complexity: O(n+m)
func intersection(nums1 []int, nums2 []int) []int {
	intersections := make(map[int]bool)
	seen2 := make(map[int]bool)

	for _, num := range nums2 {
		seen2[num] = true
	}

	for _, num := range nums1 {
		if _, ok := seen2[num]; ok {
			intersections[num] = true
		}
	}

	var ans []int
	for key, _ := range intersections {
		ans = append(ans, key)
	}

	return ans
}
