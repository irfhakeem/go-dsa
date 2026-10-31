package leetcode

import (
	"cmp"
	"slices"
)

// Time Complexity: O(n log n)
// Space Complexity: O(n)
func groupThePeople(groupSizes []int) [][]int {
	stores := make(map[int][]int)

	for i, size := range groupSizes {
		_, ok := stores[size]

		if !ok {
			stores[size] = []int{i}
		} else {
			stores[size] = append(stores[size], i)
		}
	}

	// fmt.Println(stores)

	var keys []int
	for k, _ := range stores {
		keys = append(keys, k)
	}

	slices.SortFunc(keys, func(a, b int) int {
		return cmp.Compare(a, b)
	})

	var ans [][]int
	for _, k := range keys {
		i := 0
		iter := 0
		for iter < len(stores[k])/k {
			// fmt.Println(i, stores[k][i:i+k])
			ans = append(ans, stores[k][i:i+k])
			i += k
			iter++
		}

		// fmt.Println(ans)
	}

	return ans
}
