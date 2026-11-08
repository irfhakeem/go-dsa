package leetcode

// Time Complexity: O(n)
// Space Complexity: O(n)
func smallestPalindrome(s string) string {
	n := len(s)
	if n == 1 {
		return s
	}

	alph := make([]int, 26)
	for i := range n / 2 {
		alph[s[i]-'a']++
	}

	ans := make([]byte, n)
	l := 0
	r := len(s) - 1
	for i, c := range alph {
		for c > 0 {
			ans[l] = byte(i + 'a')
			ans[r] = byte(i + 'a')
			l++
			r--
			c--
		}
	}

	if n%2 != 0 {
		mid := n / 2
		ans[mid] = s[mid]
	}

	return string(ans)
}
