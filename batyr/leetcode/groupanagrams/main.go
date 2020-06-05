package main

import "fmt"

func main() {
	var s []string
	s = append(s, "eat")
	s = append(s, "tea")
	s = append(s, "tan")
	s = append(s, "ate")
	s = append(s, "nat")
	s = append(s, "bat")

	s = append(s, "tab")
	fmt.Println(groupAnagrams(s))

	s = []string{}
	fmt.Println(s)
}
func groupAnagrams(strs []string) [][]string {
	var res [][]string
	var element []string
	// var other []string
	// element[0] = strs[0]

	for i := 0; i < len(strs); i++ {
		if i == len(strs)-1 {
			for k := range res {
				if IsAnagram(res[k][0], strs[i]) {
					return res
				}
			}
		}
		element = []string{}
		element = append(element, strs[i])
		for j := i + 1; j < len(strs); j++ {

			if IsAnagram(strs[i], strs[j]) && withoutDouble(element, strs[j]) {
				element = append(element, strs[j])
			}
		}
		res = append(res, element)
		// element = []string{}
		got := true
		for b := i + 1; b < len(strs); b++ {

			if !IsAnagram(element[0], strs[b]) {
				for k := range res {
					if IsAnagram(res[k][0], strs[b]) {
						got = false
						break
					} else {
						got = true
					}
				}
				if got {
					i = b - 1
					break
				}

			}

		}
	}

	return res

}
func withoutDouble(element []string, a string) bool {
	for _, v := range element {
		if v == a {
			return false
		}
	}
	return true
}
func IsAnagram(str1, str2 string) bool {

	if CountLen(str1) != CountLen(str2) {
		return false
	}

	newarg := make([]rune, CountLen(str2))
	j := 0
	for _, two := range str2 {
		if two != ' ' {
			newarg[j] = two
			j++
		}
	}

	count := 0
	for _, one := range str1 {
		for j, two := range newarg {
			if one == two {
				count++
				newarg[j] = '#'
				break
			}
		}
	}

	return CountLen(str2) == count

}

func CountLen(str string) int {

	counter := 0

	for _, x := range str {
		if x >= 'a' && x <= 'z' {
			counter++
		}
	}
	return counter
}
