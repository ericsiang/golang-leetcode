package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	match := false
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	replaceS := nonAlphanumericRegex.ReplaceAllString(s, "")
	if len(replaceS) == 0 || len(replaceS) == 1 {
		return true
	}

	low := strings.ToLower(replaceS)
	count := len(low) / 2
	match_c := 0

	for i := 0; i < len(low); i++ {
		j := len(low) - 1 - i
		if low[i] == low[j] {
			match_c++
			if match_c == count {
				match = true
				break
			}
		} else {
			match = false
			break
		}

	}
	return match
}
