package main

// Input: strs = ["flower","flow","flight"]
// Output: "fl"

// Input: strs = ["dog","racecar","car"]
// Output: ""

func longestCommonPrefix(strs []string) string {
	//  i  0 1 2 3
	// j   f l o w e r
	// 0   f l o w
	// 1   f l i g h t
	res := ""
	// 先取第一個的字串拿來跟其他循環比較
	for i := 0; i < len(strs[0]); i++ {
		s := strs[0][i] // f l o w e r
		for j := 1; j < len(strs); j++ {
			// i 不能超過 strs[j] 的長度 ex: len(flowe) > len(flow)
			if i == len(strs[j]) || strs[j][i] != s {
				return res
			}
		}
		res += string(s)
	}

	return res
}
