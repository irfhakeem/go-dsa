package leetcode

import "strconv"

// Time Complexity: O(n)
// Space Complexity: O(n)
func totalNumbers(digits []int) int {
	freq := make([]int, 10)
	for _, d := range digits {
		freq[d]++
	}

	store := make(map[string]bool)
	for even := 0; even <= 8; even += 2 {
		if freq[even] == 0 {
			continue
		}

		freq[even]--
		evenStr := strconv.Itoa(even)

		for f := 1; f <= 9; f++ {
			if freq[f] == 0 {
				continue
			}

			freq[f]--
			fStr := strconv.Itoa(f)

			for m := 0; m <= 9; m++ {
				if freq[m] == 0 {
					continue
				}

				mStr := strconv.Itoa(m)
				tmp := fStr + mStr + evenStr
				_, ok := store[tmp]

				if !ok {
					store[tmp] = true
					// fmt.Println(store)
				}
			}

			freq[f]++
		}

		freq[even]++
	}

	return len(store)
}
