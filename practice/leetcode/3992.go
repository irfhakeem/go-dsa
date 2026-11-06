package leetcode

// Time Complexity: O(n)
// Space Complexity: O(n)
func rearrangeString(s string, x byte, y byte) string {
	var ans []rune
	countX := 0

	for _, c := range s {
		if c == rune(x) {
			countX++
			continue
		}

		ans = append(ans, c)
	}

	for i := 0; i < countX; i++ {
		ans = append(ans, rune(x))
	}

	return string(ans)
}
