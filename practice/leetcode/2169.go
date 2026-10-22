package leetcode

func countOperations(num1 int, num2 int) int {
	if num1 == 0 || num2 == 0 {
		return 0
	}

	count := 1

	for num1 != num2 {
		count++
		if num1 < num2 {
			num2 -= num1
		} else if num1 > num2 {
			num1 -= num2
		}
	}

	return count
}
