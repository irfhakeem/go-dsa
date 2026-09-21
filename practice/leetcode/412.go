package leetcode

import "fmt"

func fizzBuzz(n int) []string {
	var ans []string

	for i := range n {
		i++
		if i%3 == 0 {
			if i%5 == 0 {
				ans = append(ans, "FizzBuzz")
			} else {
				ans = append(ans, "Fizz")
			}
		} else if i%5 == 0 {
			ans = append(ans, "Buzz")
		} else {
			idx := fmt.Sprintf("%d", i)
			ans = append(ans, idx)
		}
	}

	return ans
}
