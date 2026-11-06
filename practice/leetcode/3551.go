package leetcode

import "slices"

// Time Complexity: O(n log n)
// Space Complexity: O(n)
func minSwaps(nums []int) int {
	var store [][]int

	for i, num := range nums {
		sum := 0
		tmp := num

		for tmp > 0 {
			sum += tmp % 10
			tmp /= 10
		}

		store = append(store, []int{num, sum, i})
	}
	// fmt.Println(store)

	slices.SortFunc(store, func(a, b []int) int {
		if a[1] < b[1] {
			return -1
		} else if a[1] > b[1] {
			return 1
		} else {
			if a[0] < b[0] {
				return -1
			} else if a[0] > b[0] {
				return 1
			}
		}

		return 0
	})
	// fmt.Println(store)

	visited := make([]bool, len(store))
	count := 0

	for i := range store {
		if visited[i] || store[i][2] == i {
			continue
		}

		cycleSize := 0
		j := i

		for !visited[j] {
			visited[j] = true
			j = store[j][2]
			cycleSize++
		}

		if cycleSize > 1 {
			count += cycleSize - 1
		}
	}

	return count
}
