package leetcode

// Time Complexity: O(Num)
// Space Complexity: O(Num)
func maxSumOfSquares(num int, sum int) string {
	ans := make([]rune, num)
	for i := range ans {
		ans[i] = '0'
	}

	i := 0
	for sum > 0 && i < num {
		if sum > 9 {
			ans[i] = '9'
			sum -= 9
		} else if sum <= 9 && sum > 0 {
			ans[i] = rune(sum + '0')
			sum -= sum
		}
		// fmt.Println(ans)

		i++
	}

	if sum > 0 {
		return ""
	}

	return string(ans)
}
