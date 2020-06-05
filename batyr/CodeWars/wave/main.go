package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1]
	var word string
	countLetter := countLet(args)
	res := make([]string, countLetter)
	count := 0
	for i := 0; i < countLetter; i++ {

		for j := 0; j < len(args); j++ {
			if args[j] >= 'a' && args[j] <= 'z' && count == i {
				word += string(args[j] - 32)
				count++

			} else {

				word += string(args[j])

			}

		}
		res[i] = word
		word = ""

	}

	fmt.Println(res)
}
func countLet(str string) int {
	count := 0
	for _, v := range str {
		if v >= 'a' && v <= 'z' {
			count++
		}
	}
	return count
}
