package leetcode

import (
	"math/bits"
	"slices"
)

// Time Complexity: O(n)
// Space Complexity: O(n)
func sortByBits(arr []int) []int {
	type numBit struct {
		number   int
		totalBit int
	}

	var nbs []numBit
	for _, num := range arr {
		nb := numBit{
			number:   num,
			totalBit: bits.OnesCount(uint(num)),
		}
		nbs = append(nbs, nb)
	}

	// fmt.Println(nbs)

	slices.SortFunc(nbs, func(a, b numBit) int {
		if a.totalBit == b.totalBit && a.number > b.number {
			return 1
		} else if a.totalBit == b.totalBit && a.number < b.number {
			return -1
		} else if a.totalBit > b.totalBit {
			return 1
		} else if a.totalBit < b.totalBit {
			return -1
		} else {
			return 0
		}
	})

	// fmt.Println(nbs)

	var ans []int
	for _, nb := range nbs {
		ans = append(ans, nb.number)
	}

	return ans
}
