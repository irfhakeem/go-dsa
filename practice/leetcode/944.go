package leetcode

// Time Complexity: O(n * m)
// Space Complexity: O(1)
func minDeletionSize(strs []string) int {
	count := 0
	n := len(strs[0])

	for i := 0; i < n; i++ {
		for j := 0; j < len(strs)-1; j++ {
			if (strs[j][i] - 'a') > (strs[j+1][i] - 'a') {
				count++
				break
			}
		}
	}

	return count
}
