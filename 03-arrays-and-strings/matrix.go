package arraysandstrings

func Matrix(row, col int) [][]int {
	matrix := make([][]int, row)

	for r := range row {
		matrix[r] = make([]int, col)
		for c := range col {
			matrix[r][c] = 0
		}
	}

	return matrix
}
