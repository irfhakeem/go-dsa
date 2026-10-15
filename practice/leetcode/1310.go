package leetcode

// Time Complexity: O(n)
// Space Complexity: O(n)
func xorQueries(arr []int, queries [][]int) []int {
	nq := len(queries)
	na := len(arr)
	ans := make([]int, nq)
	pref := make([]int, na)

	pref[0] = arr[0]

	for i := 1; i < na; i++ {
		pref[i] = pref[i-1] ^ arr[i]
	}

	for i, q := range queries {
		l, r := q[0], q[1]

		ans[i] = pref[r]
		if l > 0 {
			ans[i] ^= pref[l-1]
		}
	}

	return ans
}
