package leetcode

// Time Complexity: O(n)
// Space Complexity: O(n)
func firstUniqueEven(nums []int) int {
	mp := make(map[int]int)

	for _, num := range nums {
		if num%2 == 0 {
			mp[num]++
		}
	}

	for _, num := range nums {
		if num%2 == 0 && mp[num] == 1 {
			return num
		}
	}

	return -1
}
