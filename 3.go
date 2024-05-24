package main


func lengthOfLongestSubstring(s string) int {
	l, r, res := 0, 0, 0
	checkString := make(map[byte]int)
	for r < len(s) {
		index, ok := checkString[s[r]]
		if ok && index >= l { // 检查重复字符是否在当前窗口内
			l = index + 1 // 将左指针移动到重复字符的下一个
		}
		checkString[s[r]] = r
		if r-l+1 > res { // 最大窗口长度
			res = r - l + 1
		}
		r++
	}
	return res
}
