package leetcode

// Time Complexity: O(m*n)
// Space Complexity: O(m)
func maxSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	var pref [][]int
	for _, row := range grid {
		colSum := make([]int, n)
		for j, col := range row {
			if j == 0 {
				colSum[j] = col
				continue
			}

			colSum[j] = colSum[j-1] + col
		}
		pref = append(pref, colSum)
	}
	// fmt.Println(pref)

	mx := -1
	for i := 0; i < m-2; i++ {
		for j := 0; j < n-2; j++ {
			sum := 0
			up := 0
			bot := 0

			if j > 0 {
				up = pref[i][j-1]
				bot = pref[i+2][j-1]
			}

			sum += pref[i][j+2] - up

			sum += grid[i+1][j+1]

			sum += pref[i+2][j+2] - bot

			mx = max(mx, sum)
		}
	}

	return mx
}
