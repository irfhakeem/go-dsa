package leetcode

// Time Complexity: O(n)
// Space Complexity: O(n)
func partitionLabels(s string) []int {
	store := make(map[byte]int)

	for i, c := range s {
		store[byte(c)] = i
	}

	var ans []int
	l := -1
	r := 0
	farthes := store[s[r]]

	for r < len(s) {
		farthes = max(farthes, store[s[r]])

		if r == farthes {
			ans = append(ans, r-l)
			l = r
		}

		r++
	}

	return ans
}
