package leetcode

// Time Complexity: O(n)
// Space Complexity: O(n)
func maxFreqSum(s string) int {
	mp := make(map[byte]int)
	for i, _ := range s {
		mp[s[i]]++
	}

	mxV := 0
	mxC := 0
	for k, v := range mp {
		if k == 'a' || k == 'e' || k == 'i' || k == 'u' || k == 'o' {
			mxV = max(mxV, v)
		} else {
			mxC = max(mxC, v)
		}
	}

	return mxV + mxC
}
