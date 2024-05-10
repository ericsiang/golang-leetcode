package main


// Input: s = "PAYPALISHIRING", numRows = 3
// Output: "PAHNAPLSIIGYIR"

// Input: s = "PAYPALISHIRING", numRows = 4
// Output: "PINALSIGYAHRPI"
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	slice := make([]string, numRows)
	row := 0
	res := ""
	step := 1
	for i := 0; i < len(s); i++ {
		if row == (numRows - 1) { // 到 row 的最後
			step = -1  // 逆回走
		} else if row == 0 { // 回到 row 0 
			step = 1   // 正向走
		}
		slice[row] += string(s[i])
		row += step
		
	}

	for _, v := range slice {
		res += v
	}

	return res
}
