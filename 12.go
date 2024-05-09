package main

// Input: num = 3749
// Output: "MMMDCCXLIX"

// Input: num = 58
// Output: "LVIII"

// Input: num = 1994
// Output: "MCMXCIV"

func intToRoman(num int) string {
	if num > 3999 || num < 1 {
		return ""
	}
	vales := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res := ""
	for i, v := range vales {
		for num >= v {
			num -= v
			res += romans[i]
		}
	}
	return res
}
