package leetcode

import "strconv"

// Time Complexity: O(n)
// Space Complexitu: O(n)
func smallestNumber(pattern string) string {
	stack := []int{}
	ans := ""
	currNum := 1

	for i := 0; i <= len(pattern); i++ {
		stack = append(stack, currNum)

		if i == len(pattern) || pattern[i] == 'I' {
			for i := len(stack) - 1; i >= 0; i-- {
				conv := strconv.Itoa(stack[i])
				ans += conv
			}

			stack = nil
		}

		currNum++
		// fmt.Println(stack)
	}

	if len(stack) > 0 {
		for i := len(stack) - 1; i >= 0; i-- {
			conv := strconv.Itoa(stack[i])
			ans += conv
		}
	}

	return ans
}
