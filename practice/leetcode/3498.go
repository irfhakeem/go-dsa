package leetcode

// Time Complexity: O(n)
// Space Complexity: O(1)
func reverseDegree(s string) int {
	store := make(map[rune]int)
	ch := 'z'

	for i := range 26 {
		store[ch] = i + 1
		ch = rune(int(ch) - 1)
	}

	// fmt.Println(store)

	ans := 0

	for i, c := range s {
		ans += store[c] * (i + 1)
	}

	return ans
}
