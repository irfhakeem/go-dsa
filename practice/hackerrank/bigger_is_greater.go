package hackerrank

// Time Complexity: O(n)
// Space Complexity: O(1)
func biggerIsGreater(w string) string {
	// Write your code here
	n := len(w)
	s := []byte(w)

	i := n - 2
	for i >= 0 && s[i] >= s[i+1] {
		i--
	}

	if i < 0 {
		return "no answer"
	}

	j := n - 1
	for s[j] <= s[i] {
		j--
	}

	s[i], s[j] = s[j], s[i]

	for l, r := i+1, n-1; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}

	return string(s)
}
