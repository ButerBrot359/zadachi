package main

func checkValidString(s string) bool {
	l := 0
	r := 0

	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			l--
		} else {
			l++
		}
		if l < 0 {
			return false
		}
	}
	if l == 0 {
		return true
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			r--
		} else {
			r++
		}
		if r < 0 {
			return false
		}
	}
	return true
}
