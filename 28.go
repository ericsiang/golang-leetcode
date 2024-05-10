package main

// Input: haystack = "sadbutsad", needle = "sad"
// Output: 0

// Input: haystack = "leetcode", needle = "leeto"
// Output: -1
func strStr(haystack string, needle string) int {
	Lhaystack, Lneedle := len(haystack), len(needle)
	if Lhaystack == Lneedle && haystack == needle {
		return 0
	}

	for i := 0; i < Lhaystack; i++ {
		if i+Lneedle > Lhaystack{
			return -1
		}
		if haystack[i:i+Lneedle] == needle {
			return i
		}
	}

	return -1
}
