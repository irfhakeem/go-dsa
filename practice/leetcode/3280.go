package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

// Time Complexity: O(n)
// Space Complexity: O(n)
func convertDateToBinary(date string) string {
	newDate := strings.Split(date, "-")

	for i, num := range newDate {
		numInt, _ := strconv.Atoi(num)
		newDate[i] = fmt.Sprintf("%b", numInt)
	}

	return strings.Join(newDate, "-")
}
