package leetcode

// Space Complexity: O(n^2)
// Space Complexity: O(n)
func maximumLengthSubstring(s string) int {
	mx := 0
	for i, _ := range s {
		count := 1
		mp := make(map[byte]int)
		mp[s[i]]++

		for j := i + 1; j < len(s); j++ {
			count++
			mp[s[j]]++

			if mp[s[j]] > 2 {
				clear(mp)
				break
			}

			mx = max(mx, count)
		}
	}

	return mx
}
