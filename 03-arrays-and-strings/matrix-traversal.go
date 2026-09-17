package arraysandstrings

// Time Complexity: O(n)
func MatricTraversal(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			println(val)
		}
	}
}
