package leetcode

import "strconv"

// Time Complexity: O(n x log10(m))
// Space Complexity: O(n)
func countDigitOccurrences(nums []int, digit int) int {
	count := 0
	store := make(map[int]int)
	strDigit := strconv.Itoa(digit)

	for _, num := range nums {
		strNum := strconv.Itoa(num)

		if _, ok := store[num]; ok {
			count += store[num]
		} else {
			for _, c := range strNum {
				if string(c) == strDigit {
					count++
					store[num]++
				}
			}
		}
	}

	return count
}
