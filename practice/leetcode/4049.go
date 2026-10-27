package leetcode

// Time Complexity: O(n)
// Space Complexity: O(n)
func countSpecialIntegers2(nums []int) int {
	store := make(map[int][]int)
	count := 0

	for i, num := range nums {
		_, ok := store[num]

		if !ok {
			// index, diff, total, notSpecial, isSpecial
			store[num] = []int{i, 0, 1, 0, 0}
		} else {
			store[num][2] += 1

			diff := i - store[num][0]
			store[num][0] = i

			if store[num][2] == 2 {
				store[num][1] = diff
			} else if store[num][2] >= 3 && store[num][3] != 1 && diff == store[num][1] {
				store[num][4] = 1
			} else if store[num][2] >= 3 && diff != store[num][1] {
				store[num][3] = 1
				store[num][4] = 0
			}
		}
	}

	// fmt.Println(store)

	for _, v := range store {
		if v[4] == 1 {
			count++
		}
	}

	return count
}
