package leetcode

// Time Complexity: O(n * k)
// Space Complexity: O(n)
func findAndReplacePattern(words []string, pattern string) []string {
	visitedP := make(map[rune]int)
	seqP := 1
	patternR := make([]int, len(pattern))
	for i, c := range pattern {
		_, ok := visitedP[c]

		if !ok {
			visitedP[c] = seqP
			seqP++
		}

		patternR[i] = visitedP[c]
	}
	// fmt.Println(patternR)

	var ans []string
	for _, w := range words {
		tmpW := make(map[rune]int)
		seqW := 1

		for i, c := range w {
			_, ok := tmpW[c]

			if !ok {
				tmpW[c] = seqW
				seqW++
			}

			if tmpW[c] != patternR[i] {
				break
			}

			if i == len(w)-1 {
				ans = append(ans, w)
			}
		}
	}

	return ans
}
