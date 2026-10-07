package leetcode

import "regexp"

func maxDepth(s string) int {
	re := regexp.MustCompile(`[^()]`)
	newS := re.ReplaceAllString(s, "")

	ans := 0
	count := 0

	for _, ch := range newS {
		if ch == '(' {
			count++
		} else {
			count--
		}

		ans = max(ans, count)
	}

	return ans
}
