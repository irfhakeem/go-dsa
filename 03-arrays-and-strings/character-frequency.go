package arraysandstrings

func CharacterFrequency(s string) map[rune]int {
	runes := []rune(s)

	freq := make(map[rune]int)
	for _, r := range runes {
		freq[r]++
	}
	return freq
}
