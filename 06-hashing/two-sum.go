package hashing

func TwoSum(arr []int, target int) []int {
	// Time Complexity: O(n)
	seen := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if j, ok := seen[target-arr[i]]; ok {
			return []int{j, i}
		}
		seen[arr[i]] = i
	}

	return []int{}

	// Time Complexity: O(n^2)
	// for i := 0; i < len(arr); i++ {
	// 	for j := i + 1; j < len(arr); j++ {
	// 		if arr[i]+arr[j] == target {
	// 			return []int{i, j}
	// 		}
	// 	}
	// }
	// return nil
}
