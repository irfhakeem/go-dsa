package leetcode

// Time Complexity: O(n)
// Space Complexity: O(n)
func minInsertions(s string) int {
	openStack := []rune{}
	countClose := 0
	ans := 0

	for _, ch := range s {
		if ch == '(' {
			if countClose == 1 && len(openStack) > 0 {
				countClose = max(0, countClose-1)
				openStack = openStack[:len(openStack)-1]
				ans++
			} else if countClose == 1 && len(openStack) < 1 {
				countClose = max(0, countClose-1)
				ans += 2
			}

			openStack = append(openStack, ch)
		} else {
			countClose++

			if countClose == 2 && len(openStack) == 0 {
				countClose = max(0, countClose-2)
				ans++
			} else if countClose == 2 && len(openStack) > 0 {
				countClose = max(0, countClose-2)
				openStack = openStack[:len(openStack)-1]
			}
		}
	}

	if countClose > 0 {
		if len(openStack) > 0 {
			ans += (countClose + 1) / 2
			openStack = openStack[:len(openStack)-1]
		} else {
			ans += 2
		}
	}

	// fmt.Println(countClose, openStack)

	ans += len(openStack) * 2

	return ans
}
