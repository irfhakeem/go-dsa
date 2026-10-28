package leetcode

import "slices"

// Time Complexity: O(n)
// Space Complexity: O(n)
func sortVowels3913(s string) string {
	store := make(map[rune][]int)
	runes := []rune(s)
	seq := 1

	for i, c := range runes {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			_, ok := store[c]

			if ok {
				store[c][0]++
			} else {
				store[c] = []int{1, seq}
				seq++
			}
			runes[i] = '-'
		}
	}

	if len(store) == 0 {
		return s
	}

	// fmt.Println(store, runes)

	var keys []rune
	for k := range store {
		keys = append(keys, k)
	}

	slices.SortFunc(keys, func(a, b rune) int {
		if store[a][0] > store[b][0] {
			return -1
		}
		if store[a][0] < store[b][0] {
			return 1
		}
		if store[a][0] == store[b][0] && store[a][1] > store[b][1] {
			return 1
		}
		if store[a][0] == store[b][0] && store[a][1] < store[b][1] {
			return -1
		}
		return 0
	})

	// fmt.Println(store, runes, keys)

	i := 0
	for j, c := range runes {
		if i < len(keys) {
			if store[keys[i]][0] == 0 {
				i++
			}

			if c == '-' && store[keys[i]][0] != 0 {
				runes[j] = keys[i]
				store[keys[i]][0]--
			}
		}
	}

	return string(runes)
}
