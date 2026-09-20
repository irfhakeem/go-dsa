package leetcode

func groupAnagrams(strs []string) [][]string {
	var ans [][]string
	mp := make(map[[26]int]int)
	idx := 0

	for _, s := range strs {
		var key [26]int

		for _, c := range s {
			key[c-'a']++
		}

		if groupIndex, ok := mp[key]; !ok {
			mp[key] = idx
			ans = append(ans, []string{s})
			idx++
		} else {
			ans[groupIndex] = append(ans[groupIndex], s)
		}
	}

	return ans
}
