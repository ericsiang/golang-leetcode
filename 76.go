package main

// s = "ADOBECODEBANC", t = "ABC"
func minWindow(s string, t string) string {
	tMap := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}
	tmpMap := make(map[byte]int)
	matchCount, start, length, l, r := 0, 0, len(s)+1, 0, 0
	for r < len(s) {
		//延長窗口
		right := s[r]
		r++
		if _, ok := tMap[right]; ok {
			tmpMap[right]++
			if tmpMap[right] == tMap[right] {
				matchCount++
			}
		}
		//縮短窗口
		for matchCount == len(tMap) {
			if (r - l) < length {
				start = l
				length = r - l
			}
			left := s[l]
			l++
			if _, ok := tMap[left]; ok {
				if tmpMap[left] == tMap[left] {
					matchCount--
				}
				tmpMap[left]--
			}
		}

	}
	if length == len(s)+1 {
		return ""
	}
	return s[start : start+length]
}
