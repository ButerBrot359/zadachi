package main

import "fmt"

func main() {
	str := "abcde"
	a := [][]int{{0, 1}, {1, 2}, {1, 5}, {0, 2}}
	fmt.Print(stringShift(str, a))
}

func stringShift(s string, shift [][]int) string {
	count := 0
	for v := range shift {
		if shift[v][0] == 0 {
			count = count - shift[v][1]
		} else if shift[v][0] == 1 {
			count = count + shift[v][1]
		}
	}
	// fmt.Println(count)
	newString := ""
	neg := false
	if count < 0 {
		count = -count
		neg = true
	}
	for count >= len(s) {
		count = count - len(s)
	}
	if count == len(s)-1 {
		newString = s[len(s)-count:len(s)] + s[0:len(s)-count]
		return newString
	}

	if count == 0 {
		return s
	} else if neg == false {
		newString = s[len(s)-count:len(s)] + s[0:len(s)-count]
	} else if neg == true {
		newString = s[count:len(s)] + s[0:count]
	}
	return newString
}
