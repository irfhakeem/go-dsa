package leetcode

import (
	"cmp"
	"slices"
	"strconv"
)

// Time Complexity: O(n log n)
// Space Complexity: O(1)
func minimumSum(num int) int {
	nums := []rune(strconv.Itoa(num))
	// fmt.Println(nums)

	slices.SortFunc(nums, func(a, b rune) int {
		return cmp.Compare(a, b)
	})

	num1, _ := strconv.Atoi(string(nums[0]) + string(nums[2]))
	num2, _ := strconv.Atoi(string(nums[1]) + string(nums[3]))

	return num1 + num2
}
