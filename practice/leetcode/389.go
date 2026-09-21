package leetcode

// Time Complexity: O(n), Space Complexity: O(1)
// func findTheDifference(s string, t string) byte {
// 	var keyS [26]int
// 	var keyT [26]int

// 	for _, c := range s {
// 		keyS[c-'a']++
// 	}
// 	for _, c := range t {
// 		keyT[c-'a']++
// 	}

// 	for i := 0; i < 26; i++ {
// 		if keyS[i] != keyT[i] {
// 			return byte('a' + i)
// 		}
// 	}

// 	return byte('a')
// }

// Time: O(n), Space: O(1)
func findTheDifference(s string, t string) byte {
	var res byte

	for i := 0; i < len(s); i++ {
		res ^= s[i]
	}

	for i := 0; i < len(t); i++ {
		res ^= t[i]
	}

	return res
}
