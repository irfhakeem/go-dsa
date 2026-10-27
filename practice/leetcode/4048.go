package leetcode

// Time Complexity: O(n)
// Space Complexity: O(n)
func countSpecialIntegers(nums []int) int {
	store := make(map[int][]int)
	count := 0

	for i, num := range nums {
		_, ok := store[num]

		if !ok {
			store[num] = []int{i}
		} else {
			store[num] = append(store[num], i)
		}
	}

	for _, v := range store {
		if len(v) == 3 && v[1]-v[0] == v[2]-v[1] {
			count++
		}
	}

	return count
}
