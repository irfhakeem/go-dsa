package hackerrank

func chocolateFeast(n int32, c int32, m int32) int32 {
	// Write your code here
	choco := int32(n / c)
	totalEat := int32(0)
	wraper := int32(0)

	for choco > 0 {
		if wraper%m == 0 {
			choco++
		}

		choco--
		wraper++
		totalEat++
	}

	return totalEat - 1
}
