package arraysandstrings

func ArrayManipulation(arr *[]int) *[]int {
	newArr := make([]int, len(*arr))
	copy(newArr, *arr)

	return &newArr
}
