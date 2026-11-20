package leetcode

// Time Complexity: O(n)
// Space Coplexity; O(1)
func countValidPrefixes(s string) int {
	c0 := 0
	c1 := 0
	ans := 0
	for i, _ := range s {
		if s[i] == '0' {
			c0++
		} else {
			c1++
		}

		if c0 == c1+1 || c1 == c0+1 || c1 == c0 {
			ans++
		}
	}

	return ans
}
