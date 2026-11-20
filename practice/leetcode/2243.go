package leetcode

import (
	"strconv"
	"strings"
)

// Time Complexity: O(n)
// Space Complexity: O(n)
func digitSum(s string, k int) string {
	var arr []string
	sum := 0
	for len(s) > k {
		for i, n := range s {
			sum += int(n - '0')

			if i == len(s)-1 || (i+1)%k == 0 {
				str := strconv.Itoa(sum)
				arr = append(arr, str)
				sum = 0
			}
		}

		s = strings.Join(arr, "")
		arr = arr[:0]
	}

	return s
}
