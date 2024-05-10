package main

import (
	"strings"
)

// Input: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
// Output:
// [
//
//	"This    is    an",
//	"example  of text",
//	"justification.  "
//
// ]

/*
建立左右 left right 指針 來對當前行的第一個和最後一個單詞
統計當行單詞長度 sumLen := 0 

循環确定當前行可以放多少個單詞，單詞間至少一個空格

0
left
		      2
		      right
This    is    an
* right-left 相當於是該行各單詞間至少一個空格 ex: left =0 right = 2 表示在 right 指針移到 an 了，所以該行至少有2-0個空格
right < n && sumLen+len(words[right])+right-left <= maxWidth

判斷 right 有三種情況

填空格的三種情況
1. 當前行是最後一行 單詞左對齊 且單詞間只有一個空格 行末用空格補齊剩下
2. 當前不是最後一行 且只有一個單詞 該單詞左對齊 行末用空格補齊剩下
3. 當前不是最後一行 且不只有一個單詞 
當前行單詞數為 numWords 空格數為 numSpaces
要將空格平均分配在單詞之間 因此中間數有 (numWords-1) 個
單詞間至少有 avgSpaces := numSpaces / (numWords - 1) 個空格
至於多出來的 extraSpaces := numSpaces % (numWords - 1) 個空格 應填在前 extraSpaces 個單詞之間
因此 前 extraSpaces 個單詞之間 填 extraSpaces+1 個空格 其他單詞間填 avgSpaces 個空格

s1 := strings.Join(words[left:left+extraSpaces+1], blank(avgSpaces+1)) 拼接额外加一个空格的单词
s2 := strings.Join(words[left+extraSpaces+1:right], blank(avgSpaces))         拼接其余单词
*/

func fullJustify(words []string, maxWidth int) []string {
	right := 0 //行
	n := len(words)
	res := make([]string, 0)
	for {
		
		left := right // 当前行的第一个单词在 words 的位置
		sumLen := 0 // 當前行單詞長度和
		//循环 移動 right 确定當前行可以放多少個單詞，單詞間至少一個空格
		for right < n && sumLen+len(words[right])+right-left <= maxWidth {
			sumLen += len(words[right])
			right++
		}

		// 最後一行，单词從左对齐，且单词之间应只有一个空格，在行末填充剩余空格
		if right == n {
			s := strings.Join(words[left:], " ")
			res = append(res, s+blank(maxWidth-len(s)))
			return res
		}

		numWords := right - left
		numSpaces := maxWidth - sumLen

		// 当前行只有一个单词：该单词左对齐，在行末填充剩余空格
		if numWords == 1 {
			res = append(res, words[left]+blank(numSpaces))
			continue
		}

		// 計算當行平均空格
		avgSpaces := numSpaces / (numWords - 1)
		extraSpaces := numSpaces % (numWords - 1)

		s1 := strings.Join(words[left:left+extraSpaces+1], blank(avgSpaces+1)) // 拼接额外加一个空格的单词
		s2 := strings.Join(words[left+extraSpaces+1:right], blank(avgSpaces))          // 拼接其余单词
		res = append(res, s1+blank(avgSpaces)+s2)

	}

}

// blank 返回长度为 n 的由空格组成的字符串
func blank(n int) string {
	return strings.Repeat(" ", n)
}
