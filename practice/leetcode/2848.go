package leetcode

// Time Complexity: O(n^2)
// Space Complexity: O(n)
func numberOfPoints(nums [][]int) int {
	cars := make(map[int]bool)

	for _, arr := range nums {
		for i := arr[0]; i <= arr[1]; i++ {
			cars[i] = true
		}
	}

	return len(cars)
}
