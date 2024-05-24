package main

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) > len(t) {
		return false
	}
	i := 0 // s的指針
	j := 0 // t的指針
	for j < len(t) {
		if s[i]  == t[j] {
			i++
			if i >= len(s){
				return true
			}
		}

		j++
	}

	return false
}
