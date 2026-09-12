package main

import (
	"fmt"
	"regexp"
	"strings"
)

// CHARS
// .       any character
// \d      digit
// \w      word character
// \s      whitespace

// CLASS
// [abc]       a/b/c
// [a-z]       lowercase
// [A-Z]       uppercase
// [0-9]       digit
// [^0-9]      not digit

// QUANTIFIER
// *           0+
// +           1+
// ?           0/1
// {n}         exactly n
// {n,m}       n to m
// {n,}        n+

// ANCHOR
// ^           beginning / not when inside a character class [^]
// $           end

// GROUP
// (...)       capture group
// (?:...)     non-capturing group
// |           OR

// ESCAPE
// \.          literal .
// \+          literal +
// \(          literal (

func main() {
	exampleStr1 := "Hello, My name is Hakim. Nice to meet you!"
	exampleStr2 := "110001"
	exampleStr3 := "  John_Doe @2026! "

	// remove whitespace and symbols, then turn all uppercase to lowercase
	str1Regex := regexp.MustCompile(`[^a-zA-Z]`)
	resultStr1 := str1Regex.ReplaceAllString(exampleStr1, "")
	fmt.Println(strings.ToLower(resultStr1))

	// binary must start with 1, have 0s and 1s only, have even 0s, have at least 2 0s
	str2Regex := regexp.MustCompile(`^1(1*01*0)+$`)
	matched := str2Regex.MatchString(exampleStr2)
	fmt.Println(matched)

	// turn string into valid username (lowercase, symbol allowed only underscore, no whitespace)
	str3Regex := regexp.MustCompile(`[^a-zA-Z0-9_]`)
	resultStr3 := str3Regex.ReplaceAllString(exampleStr3, "")
	fmt.Println(strings.ToLower(resultStr3))
}
