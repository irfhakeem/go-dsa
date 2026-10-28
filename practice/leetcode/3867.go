package leetcode

import (
	"cmp"
	"slices"
)

// Time Complexity: O(n^m)
// Space Complexity: O(n)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdSum(nums []int) int64 {
	mx := -1
	var prefGcd []int

	for _, num := range nums {
		mx = max(mx, num)
		prefGcd = append(prefGcd, gcd(num, mx))
	}

	slices.SortFunc(prefGcd, func(a, b int) int {
		return cmp.Compare(a, b)
	})

	// fmt.Println(prefGcd)

	ans := int64(0)
	l := 0
	r := len(prefGcd) - 1

	for l < r {
		ans += int64(gcd(prefGcd[l], prefGcd[r]))
		l++
		r--
	}

	return ans
}
