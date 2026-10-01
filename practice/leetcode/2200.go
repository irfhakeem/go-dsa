package leetcode

// Time Complexity: O(n * k)
// Space Complexity: O(n + k)
func findKDistantIndices(nums []int, key int, k int) []int {
	var keys []int

	for i, num := range nums {
		if num == key {
			keys = append(keys, i)
		}
	}

	var result []int
	seen := make(map[int]bool)

	for _, idx := range keys {
		for i := idx - k; i <= idx+k; i++ {
			_, ok := seen[i]

			if i >= 0 && i < len(nums) && !ok {
				result = append(result, i)
				seen[i] = true
			}
		}
	}
	return result
}
