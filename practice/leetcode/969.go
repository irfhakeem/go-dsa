package leetcode

import "slices"

// Time Complexity: O(n^2)
// Space Complexity: O(n)
func pancakeSort(arr []int) []int {
	var ans []int
	end := len(arr)
	mx := end

	for end > 1 {
		var maxIdx int

		for i := 0; i < end; i++ {
			if arr[i] == mx {
				maxIdx = i
				break
			}
		}

		if maxIdx != 0 {
			slices.Reverse(arr[:maxIdx+1])
			ans = append(ans, maxIdx+1)
		}

		slices.Reverse(arr[:end])
		ans = append(ans, end)

		mx--
		end--
	}

	return ans
}
