package leetcode

// Time Complexity: O(n)
// Space Complexity: O(n*m) m = total freq
func topKFrequent(nums []int, k int) []int {
	store := make(map[int]int)

	for _, num := range nums {
		store[num]++
	}

	mx := -1
	for _, v := range store {
		mx = max(mx, v)
	}

	// fmt.Println(mx)

	rank := make([][]int, mx)
	for k, v := range store {
		rank[v-1] = append(rank[v-1], k)
	}

	// fmt.Println(rank)

	var ans []int
	for i := len(rank) - 1; i >= 0; i-- {
		ans = append(ans, rank[i]...)

		if len(ans) >= k {
			return ans[:k]
		}
	}

	return ans
}
