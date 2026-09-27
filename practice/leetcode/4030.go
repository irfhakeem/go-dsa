package leetcode

import "fmt"

func isPalindromic(s string) bool {
	var binStr string

	for _, c := range s {
		binStr += fmt.Sprintf("%08b", c)
	}

	bin := []byte(binStr)

	l := 0
	r := len(bin) - 1

	for l < r {
		if string(bin[l]) != string(bin[r]) {
			return false
		}

		r--
		l++
	}

	return true
}
