package main

import "fmt"

func main() {
	a := 19
	fmt.Println(isHappy(a))

}

func isHappy(n int) bool {
	if n == 0 {
		return false
	} else if n == 1 {
		return true
	}
	number := n
	count := 0
	for number > 1 {
		if count == 10 {
			return false
		}
		strnum := Itoa(number)
		number = 0
		for _, v := range strnum {

			power2 := Atoi(string(v)) * Atoi(string(v))
			number = number + power2
		}
		count++
	}
	return true
}
func Atoi(s string) int {
	min := 0
	plu := 0
	sum := 0
	for i, v := range s {
		if s[i] >= '0' && s[i] <= '9' {
			num := 0
			for j := '0'; j < v; j++ {
				num++
			}
			sum = sum*10 + num
		} else if s[i] == '-' && i == 0 {
			min++
		} else if s[i] == '+' && i == 0 {
			plu++
		} else {
			return 0
		}
	}
	if min == 1 {
		sum = -sum
	}
	if min != 1 && sum < 0 {
		return 0
	}
	if min == 1 && sum > 0 {
		return 0
	}
	return sum

}
func Itoa(n int) string { // better
	res := ""
	t := 1
	if n < 0 {
		t = -1
	}
	n = n * t

	for n > 0 {
		res = string(n%10+48) + res
		n /= 10
	}
	if t == -1 {
		res = "-" + res
	}
	return res
}
