package leetcode

// Time Complexity: O(n)
// Space Complexity: O(n)
func stringHash(s string, k int) string {
	i := 0
	iter := 0
	sum := 0

	var ans []rune

	for i < len(s) {
		iter++
		sum += int(s[i]) - 97

		if iter == k {
			// fmt.Println(sum%26)
			ans = append(ans, rune((sum%26)+97))
			sum = 0
			iter = 0
		}

		i++
	}

	if sum > 0 {
		// fmt.Println(sum%26)
		ans = append(ans, rune((sum%26)+97))
	}

	return string(ans)
}
