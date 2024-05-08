package main

// Input: s = "III"
// Output: 3

// Input: s = "LVIII"
// Output: 58

// Input: s = "MCMXCIV"
// Output: 1994

func romanToInt(s string) int {
	romanMap := make(map[string]int)
	romanMap["I"] = 1
	romanMap["V"] = 5
	romanMap["X"] = 10
	romanMap["L"] = 50
	romanMap["C"] = 100
	romanMap["D"] = 500
	romanMap["M"] = 1000
	total := 0
	for i := 0; i < len(s); i++ {
		if i > len(s)-1 {
			break
		}

		if string(s[i]) == "I" {
			if i+1 <= len(s)-1 && string(s[i+1]) == "V" {
				total += 4
				i++
				continue
			}

			if i+1 <= len(s)-1 && string(s[i+1]) == "X" {
				total += 9
				i++
				continue
			}
		}

		if string(s[i]) == "X" {
			if i+1 <= len(s)-1 && string(s[i+1]) == "L" {
				total += 40
				i++
				continue
			}

			if i+1 <= len(s)-1 && string(s[i+1]) == "C" {
				total += 90
				i++
				continue
			}
		}

		if string(s[i]) == "C" {
			if i+1 <= len(s)-1 && string(s[i+1]) == "D" {
				total += 400
				i++
				continue
			}

			if i+1 <= len(s)-1 && string(s[i+1]) == "M" {
				total += 900
				i++
				continue
			}
		}

		if val, ok := romanMap[string(s[i])]; ok {
			total += val
		}
	}

	return total
}
