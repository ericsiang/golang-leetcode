package main

// Input: s = "Hello World"
// Output: 5

// Input: s = "   fly me   to   the moon  "
// Output: 4
func lengthOfLastWord(s string) int {
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			count++
		} else if s[i] == ' ' && count != 0 {
			break
		}
	}
	return count
}
