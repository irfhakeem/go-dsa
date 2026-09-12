package leetcode

func isAnagram(s string, t string) bool {
	runeS := []rune(s)
	runeT := []rune(t)
	n := len(runeS)
	m := len(runeT)

	if n != m {
		return false
	}

	mapS := make(map[rune]int)
	mapT := make(map[rune]int)

	for i := 0; i < n; i++ {
		mapS[runeS[i]]++
		mapT[runeT[i]]++
	}

	for r, count := range mapT {
		_, exist := mapS[r]
		if !exist {
			return false
		}

		if count != mapS[r] {
			return false
		}
	}

	return true
}
