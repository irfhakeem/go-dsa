package arraysandstrings

// Time Complexity: O(n)
func Transpose(matrix [][]int) [][]int {
	var newMatrix [][]int

	colLen := len(matrix[0])
	rowLen := len(matrix)

	for i := 0; i < colLen; i++ {
		newRow := make([]int, rowLen)
		for j := 0; j < rowLen; j++ {
			newRow[j] = matrix[j][i]
		}
		newMatrix = append(newMatrix, newRow)
	}
	return newMatrix
}
