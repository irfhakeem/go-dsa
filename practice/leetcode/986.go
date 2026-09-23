package leetcode

// Time Complexity: O(n + m)
// Space Complexity: O(n)
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var res [][]int

	a := 0
	b := 0

	for a < len(firstList) && b < len(secondList) {
		var tmp []int

		startA := firstList[a][0]
		startB := secondList[b][0]

		endA := firstList[a][1]
		endB := secondList[b][1]

		tmp = append(tmp, max(startA, startB), min(endA, endB))

		if tmp[1] >= tmp[0] {
			res = append(res, tmp)
		}

		if endA < endB {
			a++
		} else if endA > endB {
			b++
		} else {
			a++
			b++
		}
	}

	return res
}
