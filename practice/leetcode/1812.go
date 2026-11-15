package leetcode

import "strconv"

// Time Complexity: O(n)
// Space Complexity: O(1)
func squareIsWhite(coordinates string) bool {
	alpInt := coordinates[0] - 'a'
	numInt, _ := strconv.Atoi(string(coordinates[1]))

	if alpInt%2 == 0 && numInt%2 == 0 {
		return true
	} else if alpInt%2 == 0 && numInt%2 != 0 {
		return false
	} else if alpInt%2 != 0 && numInt%2 == 0 {
		return false
	} else {
		return true
	}
}
