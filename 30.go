package main


// barfoothefoobarman
// [ "foo", "bar" ]
func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}

	res := []int{}
	wordsMap := make(map[string]int)
	for _, word := range words {
		wordsMap[word]++
	}
	w, length := len(words[0]), len(words)*len(words[0])
	for start := 0; start+length <= len(s); start++ {
		tmpMap := make(map[string]int)
		//遍歷這個窗口的詞 檢查是否在 words 中出現過
		for j := 0; j < length; j += w {
			tmp := s[start+j : start+j+w]
			// 如果 words 中包含這個詞
			if _, ok := wordsMap[tmp]; ok {
				tmpMap[tmp]++
				// 判斷 tmp 在當前窗口中出現的次數，和 words 中的次數比對
				// 如果 tmpMap 超過 words 中的次數，表示不匹配，直接跳出
				if tmpMap[tmp] > wordsMap[tmp] {
					break
				}
			} else {
				// words 中不包含這個詞，直接跳出
				break
			}

			//	如果窗口中的單詞都匹配上了
			if j == length-w {
				res = append(res, start)
			}
		}

	}

	return res
}
