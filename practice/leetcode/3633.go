package leetcode

// Time Complexity: O(n^m)
// Space Complexity: O(1)
func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	mn := 999999999

	for i, landStart := range landStartTime {
		landF := landStart + landDuration[i]

		for j, waterStart := range waterStartTime {
			f := max(landF, waterStart) + waterDuration[j]
			mn = min(mn, f)
		}
	}

	for i, waterStart := range waterStartTime {
		waterF := waterStart + waterDuration[i]

		for j, landStart := range landStartTime {
			f := max(waterF, landStart) + landDuration[j]
			mn = min(mn, f)
		}
	}

	return mn
}
