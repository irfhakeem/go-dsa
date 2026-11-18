package leetcode

// Time Complexity: O(n x m)
// Space Complexity: O(n)
func removeOccurrences(s string, part string) string {
	sb := []byte(s)

	var ans []byte
	for len(sb) > 0 {
		ans = append(ans, sb[0])
		sb = sb[1:]

		if len(ans) >= len(part) {
			if string(ans[len(ans)-len(part):]) == part {
				ans = ans[:len(ans)-len(part)]
			}
		}

		// fmt.Println(sb, ans)
	}

	return string(ans)
}
