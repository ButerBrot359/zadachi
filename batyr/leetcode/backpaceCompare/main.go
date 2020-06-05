package main

import "fmt"

func main() {
	S := "a##c"
	T := "#a#c"
	fmt.Print(backspaceCompare(S, T))
}
func backspaceCompare1(S string, T string) bool {
	if Compare(S) == Compare(T) {
		return true
	} else {
		return false
	}
}
func Compare(str string) string {

	s := ""
	back := 0
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '#' {
			back++
		} else if str[i] >= 'a' && str[i] <= 'z' && back != 0 {
			back--
		} else {
			s = string(str[i]) + s
		}
	}
	return s
}
