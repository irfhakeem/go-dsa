package leetcode

import "slices"

// Time Complexity: O(n^2)
// Space Complexity: O(n)
func rearrangeArray(nums []int) []int {
	slices.Sort(nums)

	var uNum []int
	store := make(map[int]int)
	for _, num := range nums {
		_, ok := store[num]

		if !ok {
			uNum = append(uNum, num)
		}

		store[num]++
	}
	l
	// fmt.Println(uNum, store)

	var ans []int
	i := 0
	for len(uNum) > 0 {
		ans = append(ans, uNum[i])
		store[uNum[i]]--

		if store[uNum[i]] == 0 {
			tmp := i

			if i >= len(uNum)-1 {
				i = 0
			}

			uNum = slices.Delete(uNum, tmp, tmp+1)

			continue
		}

		if i >= len(uNum)-1 {
			i = 0
			continue
		}

		i++
	}

	return ans
}
