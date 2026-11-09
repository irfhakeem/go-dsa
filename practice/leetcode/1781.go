package leetcode

// Time Complexity: O(n^2)
// Space Complexity: O(1)
func beautySum(s string) int {
	count := 0

	for l := 0; l < len(s); l++ {
		arr := make([]int, 26)

		for r := l; r < len(s); r++ {
			arr[s[r]-'a']++

			mx, mn := 0, len(s)

			for _, v := range arr {
				if v > 0 {
					mx = max(mx, v)
					mn = min(mn, v)
				}
			}

			count += mx - mn
			// fmt.Println(arr, count, mx, mn)
		}
	}

	return count
}
