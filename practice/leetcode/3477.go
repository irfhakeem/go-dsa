package leetcode

// Time Complexity: O(n^2)
// Space Complexity: O(1)
func numOfUnplacedFruits(fruits []int, baskets []int) int {
	count := 0
	for i, _ := range fruits {
		haveBasket := false
		for j := 0; j < len(baskets); j++ {
			if fruits[i] <= baskets[j] && baskets[j] != 0 {
				baskets[j] = 0
				haveBasket = true
				break
			}
		}

		if !haveBasket {
			count++
		}
	}

	return count
}
