package leetcode

// Time Complexity: O(n)
// Space Complexity: O(1) - Because max freq is 26 alphabet
func minSteps(s string, t string) int {
	freq := make(map[byte]int)
	for i := range len(s) {
		freq[s[i]]--
		freq[t[i]]++
	}

	count := 0
	for _, v := range freq {
		if v < 0 {
			count += v
		}
	}

	return count * (-1)
}
