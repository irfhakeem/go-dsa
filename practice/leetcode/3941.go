package leetcode

// Time Complexity: O(n)
// Space Complexity: O(n)
func passwordStrength(password string) int {
	ans := 0
	mp := make(map[rune]bool)
	for _, c := range password {
		_, ok := mp[c]

		if !ok {
			mp[c] = true

			if c >= 'a' && c <= 'z' {
				ans += 1
			} else if c >= 'A' && c <= 'Z' {
				ans += 2
			} else if c >= '0' && c <= '9' {
				ans += 3
			} else {
				ans += 5
			}
		}
	}

	return ans
}
