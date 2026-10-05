package hackerrank

// Time Complexity : O(n)
// Space Complexity : O(n)
func appendAndDelete(s string, t string, k int32) string {
	// Write your code here
	sArr := []rune(s)
	tArr := []rune(t)

	nS := int32(len(sArr))
	nT := int32(len(tArr))

	prefix := int32(0)

	for prefix < nS && prefix < nT &&
		sArr[prefix] == tArr[prefix] {
		prefix++
	}

	minOps := (nS - prefix) + (nT - prefix)

	if minOps > k {
		return "No"
	}

	remaining := k - minOps
	// fmt.Println(minOps, remaining)

	if remaining%2 == 0 {
		return "Yes"
	}

	if k >= nS+nT {
		return "Yes"
	}

	return "No"
}
