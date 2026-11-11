package leetcode

import "math"

// Time Complexity: O(n)
// Space Complexity: O(1)
// func pivotInteger(n int) int {
// 	if n == 1 {
// 		return 1
// 	}
//
// 	totalSum := 0
// 	for i := range n {
// 		totalSum += i + 1
// 	}
//
// 	leftSum := 0
// 	for i := range n {
// 		rightSum := totalSum - i - leftSum
// 		// fmt.Println(leftSum, rightSum)
//
// 		if rightSum == leftSum {
// 			return i
// 		}
//
// 		leftSum += i
// 	}
//
// 	return -1
// }

// Time Complexity: O(1)
// Spcae Complexity: O(1)
func pivotInteger(n int) int {
	total := n * (n + 1) / 2
	pivot := int(math.Sqrt(float64(total)))

	if pivot*pivot == total {
		return pivot
	}

	return -1
}
