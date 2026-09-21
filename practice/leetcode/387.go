package leetcode

func firstUniqChar(s string) int {
	type CharInfo struct {
		Index int
		Count int
	}

	mp := make(map[rune]CharInfo)
	runes := []rune(s)

	for i, c := range runes {
		if val, ok := mp[c]; !ok {
			info := CharInfo{
				Index: i,
				Count: 1,
			}
			mp[c] = info
		} else {
			val.Count++
			mp[c] = val
		}
	}

	for i, c := range runes {
		if mp[c].Count == 1 {
			return i
		}
	}

	return -1
}
