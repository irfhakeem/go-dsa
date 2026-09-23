package leetcode

import "unicode"

func reverseOnlyLetters(s string) string {
	bytes := []byte(s)
	var q []byte
	var idxs []int

	for i, b := range bytes {
		if unicode.IsLetter(rune(b)) {
			q = append(q, b)
			idxs = append(idxs, i)
		}
	}

	l := 0
	r := len(idxs) - 1

	for r > l {
		tmp := bytes[idxs[l]]
		bytes[idxs[l]] = bytes[idxs[r]]
		bytes[idxs[r]] = tmp
		l++
		r--
	}

	return string(bytes)
}
