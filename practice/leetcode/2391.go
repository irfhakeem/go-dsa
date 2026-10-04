package leetcode

// Time Complexity: O(n + m)
// Space Complexity: O(n)
func garbageCollection(garbage []string, travel []int) int {
	glass := 0
	metal := 0
	paper := 0
	totalTravel := 0

	lastGI := -1
	lastMI := -1
	lastPI := -1

	var p []int

	for i, g := range garbage {
		if i != 0 {
			totalTravel += travel[i-1]
		}
		p = append(p, totalTravel)

		for _, c := range g {
			switch c {
			case 'M':
				metal++
				lastMI = i
			case 'P':
				paper++
				lastPI = i
			case 'G':
				glass++
				lastGI = i
			}
		}
	}

	if lastMI != -1 {
		metal += p[lastMI]
	}
	if lastPI != -1 {
		paper += p[lastPI]
	}
	if lastGI != -1 {
		glass += p[lastGI]
	}

	return metal + glass + paper
}
