package leetcode

import "math"

// Time Complexity: O(n^2)
// Space Complexity: O(n)
// func minOperations(boxes string) []int {
// 	var ans []int
//
// 	for i, _ := range boxes {
// 		count := 0
//
// 		for j, b := range boxes {
// 			if i == j {
// 				continue
// 			} else if b == '1' {
// 				count += int(math.Abs(float64(j) - float64(i)))
// 			}
// 		}
//
// 		ans = append(ans, count)
// 	}
//
// 	return ans
// }

// Time Complexity: O(n)
// Space Complexity: O(n)
func minOperations(boxes string) []int {
	var p []int

	count := 0
	counter := 0

	for _, b := range boxes {
		count += counter
		if b == '1' {
			counter++
		}

		p = append(p, count)
	}

	ans := make([]int, len(boxes))
	count = 0
	counter = 0

	for i := len(boxes) - 1; i >= 0; i-- {
		count += counter
		if boxes[i] == '1' {
			counter++
		}

		ans[i] = int(math.Abs(float64(p[i]) + float64(count)))
	}

	return ans
}
