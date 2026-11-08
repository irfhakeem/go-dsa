package leetcode

import "strings"

// Time Complexity: O(n * l)
// Space Complexity: O(n * l)
func numOfPairs(nums []string, target string) int {
	mp := make(map[string]string)
	counting := make(map[string]int)
	for _, num := range nums {
		pref := strings.TrimSuffix(target, num)
		if len(pref) < len(target) {
			mp[pref] = num
		}

		counting[num]++
	}
	// fmt.Println(mp, counting)

	count := 0
	for k, v := range mp {
		if k == v {
			n := counting[k]
			count += n * (n - 1)
		} else {
			count += counting[k] * counting[v]
		}
	}

	return count
}
