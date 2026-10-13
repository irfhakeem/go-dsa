package leetcode

import "sort"

// Time Complexity: O(n log n)
// Space Complexity: O(n)
func sortVowels(s string) string {
	vowels := map[rune]int{
		'A': 1,
		'E': 2,
		'I': 3,
		'O': 4,
		'U': 5,
		'a': 6,
		'e': 7,
		'i': 8,
		'o': 9,
		'u': 10,
	}

	storeV := []rune{}
	sArr := []rune(s)

	for _, ch := range sArr {
		if _, ok := vowels[ch]; ok {
			storeV = append(storeV, ch)
		}
	}

	sort.Slice(storeV, func(i, j int) bool {
		return vowels[storeV[i]] < vowels[storeV[j]]
	})

	vI := 0
	for i, ch := range sArr {
		if _, ok := vowels[ch]; ok {
			sArr[i] = storeV[vI]
			vI++
		}
	}

	return string(sArr)
}
