package hackerrank

import "math"

// Time Complexity: O(n^2)
// Space Complexity: O(n)
func minimumDistances(a []int32) int32 {
	// Write your code here
	seen := make(map[int32]bool)
	ans := int32(99999)

	for l, num := range a {
		_, ok := seen[num]
		if ok {
			continue
		}

		for r := l + 1; r < len(a); r++ {
			if a[l] == a[r] {
				seen[num] = true

				// fmt.Println(l, r)
				ans = min(ans, int32(math.Abs(float64(l)-float64(r))))
				break
			}
		}
	}

	if ans == int32(99999) {
		return int32(-1)
	}

	return ans
}
