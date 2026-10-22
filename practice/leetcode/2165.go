package leetcode

import (
	"cmp"
	"slices"
	"strconv"
)

// Time Complexity: O(d log d)
// Space Complexity: O(d)
func smallestNumber(num int64) int64 {
	isNeg := false

	if num < 0 {
		isNeg = true
	}

	sNum := strconv.FormatInt(num, 10)
	rNums := []rune(sNum)
	// fmt.Println(rNums)

	if len(rNums) < 2 {
		return num
	}

	if isNeg {
		rNums = rNums[1:]

		slices.SortFunc(rNums, func(a, b rune) int {
			return cmp.Compare(b, a)
		})
	} else {
		slices.SortFunc(rNums, func(a, b rune) int {
			return cmp.Compare(a, b)
		})

		preffZero := 0

		for _, num := range rNums {
			if num == '0' {
				preffZero++
			} else {
				break
			}
		}

		if preffZero > 0 {
			tmp := rNums[0]
			rNums[0] = rNums[preffZero]
			rNums[preffZero] = tmp

			// fmt.Println(string(rNums))
		}
	}

	if isNeg {
		sNum = "-" + string(rNums)
	} else {
		sNum = string(rNums)
	}

	ans, _ := strconv.ParseInt(sNum, 10, 64)

	return ans
}
