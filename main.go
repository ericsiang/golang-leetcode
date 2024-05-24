package main

import "fmt"

func main() {
	fmt.Println("---------leetcode 150-----------")
	//1.
	// nums := []int{3,2,4}
	// fmt.Println(twoSum(nums,6))

	//3
	// s3 := "abcabcbb"
	// s3 = "pwwkew"
	// s3 = "cdd"
	// s3 = "bpfbhmipx"
	// r := lengthOfLongestSubstring(s3)
	// fmt.Println(r)

	//
	// s3 := "PAYPALISHIRING"
	// numRows3 := 3
	// s3 = "PAYPALISHIRING"
	// numRows3 = 4
	// s :=convert(s3,numRows3)
	// fmt.Println(s)

	//11
	// nums11 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	// r := maxArea(nums11)
	// fmt.Println(r)

	//12
	// num12 := 3749
	// num12 := 58
	// num12 := 1994
	// roman :=intToRoman(num12)
	// fmt.Println(roman)

	//13
	// s := "III"
	// s := "LVIII"
	// s := "MCMXCIV"
	// c :=romanToInt(s)
	// fmt.Println(c)

	//14
	// strs14 := []string{"flower","flow","flowht"}
	// strs14 := []string{"dog","racecar","car"}
	// s :=longestCommonPrefix(strs14)
	// fmt.Println(s)

	//15
	// nums15 := []int{-1, 0, 1, 2, -1, -4}
	// r := threeSum(nums15)
	// fmt.Println(r)

	//26.
	// nums26 := []int{1, 1, 2}
	// fmt.Println("nums26 : ", nums26)
	// k := removeDuplicates(nums26)
	// fmt.Println("k : ", k)

	//27.
	// nums27 := []int{3,2,2,3}
	// nums27_1 := []int{0,1,2,2,3,0,4,2}
	// removeElement(nums27,2)
	// fmt.Println(nums27)

	//28
	// haystack := "sadbutsad"
	// needle := "sad"
	// haystack = "leetcode"
	// needle = "leeto"
	// haystack = "hello"
	// needle = "ll"
	// n := strStr(haystack, needle)
	// fmt.Println(n)

	//30
	// s30 := "barfoothefoobarman"
	// words := []string{"foo", "bar"}
	// r := findSubstring(s30, words)
	// fmt.Println(r)

	//42
	// nums42 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// res := trap(nums42)
	// fmt.Println(res)

	//45
	// nums45 := []int{2, 3, 1, 1, 4}
	// k := jump(nums45)
	// fmt.Println("k:", k)

	//58
	// s58 := "Hello World"
	// s58 := "   fly me   to   the moon  "
	// leng := lengthOfLastWord(s58)
	// fmt.Println(leng)

	//68
	// words58 := []string{"This", "is", "an", "example", "of", "text", "justification."}
	// maxWidth58 := 16
	// res := fullJustify(words58, maxWidth58)
	// fmt.Println(res)
	// for i, v := range res {
	// 	fmt.Println(i, v)

	// }

	// 76
	s := "ADOBECODEBANC"
	t := "ABC"
	r := minWindow(s, t)
	fmt.Println(r)


	//80
	// nums80 := []int{1,1,1,2,2,3}
	// k := removeDuplicates2(nums80)
	// fmt.Println("k : ", k)

	//88.
	// nums88 := []int{1, 2, 3, 0, 0, 0}
	// m := 3
	// nums88_1 := []int{2, 5, 6}
	// n := 3
	// merge(nums88, m, nums88_1, n)
	// fmt.Println(nums88)

	//125
	// s125 := "A man, a plan, a canal: Panama"
	// s125 = "race a car"
	// r := isPalindrome(s125)
	// fmt.Println(r)

	//135
	// nums135 := []int{1,2,2}
	// t :=candy(nums135)
	// fmt.Println(t)

	//151
	// s151 := "the sky is blue"
	// s151 := "  hello world  "
	// newS :=reverseWords(s151)
	// fmt.Println(newS)

	//167
	// nums167 := []int{2, 7, 11, 15}
	// r :=twoSum2(nums167, 9)
	// fmt.Println(r)

	//189
	// nums189 := []int{1,2,3,4,5,6,7}
	// rotate(nums189,3)
	// fmt.Println(nums189)

	//209
	// nums209 := []int{2, 3, 1, 2, 4, 3}
	// r := minSubArrayLen(7, nums209)
	// fmt.Println(r)

	//217.
	// nums := []int{1, 2, 3, 1}
	// fmt.Println(containsDuplicate(nums))

	// 238
	// nums238:=[]int{1,2,3,4}
	// productExceptSelf(nums238)

	//274
	// nums274:=[]int{3,0,6,1,5}
	// hIndex(nums274)

	// 380
	// s := Constructor()
	// b := s.Insert(1)
	// b2 := s.Remove(2)
	// b3 := s.Insert(2)
	// r := s.GetRandom()
	// b4 := s.Remove(1)
	// b5 := s.Insert(2)
	// r2 := s.GetRandom()

	// fmt.Printf("[null,%b,%b,%b,%d,%b,%b,%d]",b,b2,b3,r,b4,b5,r2)
	// fmt.Println()

	// 392
	// s := "abc"
	// t := "awbdtdc"
	// r := isSubsequence(s, t)
	// fmt.Println(r)
}
