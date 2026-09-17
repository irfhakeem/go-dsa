package searching

func LinearSearch(arr []int, target int) int {
	for _, val := range arr {
		if val == target {
			return val
		}
	}
	return -1
}
