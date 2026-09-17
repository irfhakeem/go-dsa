package arraysandstrings

import (
	"regexp"
	"strings"
)

func SplitString(s string, delimiter string) []string {
	re := regexp.MustCompile(`\s`)
	s = re.ReplaceAllString(s, delimiter)

	newString := strings.Split(s, delimiter)
	return newString
}

func JoinString(s []string, delimiter string) string {
	newString := strings.Join(s, delimiter)
	return newString
}
