package leetcode

// Time Complexity: O(n^3)
// Space Complexity: O(1)
// func isPalindrome(s string) bool {
// 	l := 0
// 	r := len(s) - 1
//
// 	for l < r {
// 		if s[l] != s[r] {
// 			return false
// 		}
// 		l++
// 		r--
// 	}
//
// 	return true
// }
//
// func countSubstrings(s string) int {
// 	count := 0
// 	start := 0
//     for start < len(s) {
//         end := len(s)
//
//         for end > start {
//             if isPalindrome(s[start:end]) {
//                 count++
//             }
//             end--
//         }
//
//         start++
//     }
//
// 	return count
// }

// Time Complexity: O(n^2)
// Space Complexity: O(1)
func countSubstrings(s string) int {
	count := 0

	for i := 0; i < len(s); i++ {
		l := i
		r := i

		for l >= 0 && r < len(s) && s[l] == s[r] {
			count++
			l--
			r++
		}

		l = i
		r = i + 1

		for l >= 0 && r < len(s) && s[l] == s[r] {
			count++
			l--
			r++
		}
	}

	return count
}
