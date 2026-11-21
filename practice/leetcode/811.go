package leetcode

import (
	"strconv"
	"strings"
)

// Time Complexity: O(n x l)
// Space Complexity: O(n x l)
func subdomainVisits(cpdomains []string) []string {
	mp := make(map[string]int)

	for _, domain := range cpdomains {
		res := strings.Split(domain, " ")
		ts, ds := res[0], res[1]

		t, _ := strconv.Atoi(ts)
		mp[ds] += t

		for i := 0; i < len(ds); i++ {
			if ds[i] == '.' {
				mp[ds[i+1:]] += t
			}
		}
	}

	var ans []string

	for k, v := range mp {
		var sb strings.Builder

		vs := strconv.Itoa(v)
		sb.WriteString(vs)
		sb.WriteString(" ")
		sb.WriteString(k)

		ans = append(ans, sb.String())
	}

	return ans
}
