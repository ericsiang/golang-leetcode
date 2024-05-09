package main

import (
	"strings"
)

// Input: s = "the sky is blue"
// Output: "blue is sky the"

// Input: s = "  hello world  "
// Output: "world hello"

// Input: s = "a good   example"
// Output: "example good a"
func reverseWords(s string) string {
	sArr := strings.Split(s, " ")
	var newArr []string
	for i := len(sArr) - 1; i >= 0; i-- {
		// fmt.Println(sArr[i])
		if sArr[i] != "" {
			newArr = append(newArr, sArr[i])
		}
	}
	// fmt.Println(len(newArr))
	return strings.Join(newArr, " ")
}
