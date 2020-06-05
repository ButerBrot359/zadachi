// 1 res
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isHappy0(n int) bool {
	mapChecker := map[int]bool{n: true}
	for {
		n = calc(n)
		if n == 1 {
			break
		}
		if mapChecker[n] {
			return false
		} else {
			mapChecker[n] = true
		}
	}
	return true
}

func calc(n int) int {
	res := 0
	for {
		l := n % 10
		res = res + l*l
		n = n / 10
		if n == 0 {
			break
		}
	}
	return res
}

// ________________________________
// 2 res
func isHappy1(n int) bool {
	m := make(map[int]bool)
	for n > 1 {
		m[n] = true
		var x int
		for x, n = n, 0; x > 0; x /= 10 {
			d := x % 10
			n += d * d
		}
		if m[n] {
			return false
		}
	}
	return true
}

// _________________________________
// 3 res
func isHappy2(n int) bool {
	if n == 4 {
		return false
	}
	b := 0
	for _, value := range strings.Split(strconv.Itoa(n), "") {
		c, _ := strconv.Atoi(value)
		b += c * c
	}
	if b == 1 {
		return true
	}
	return isHappy(b)
}

// ___________________
// 4res
func isHappy3(n int) bool {
	sum := 0
	tmp := 0
	for {
		fmt.Println("n:", n)
		sum += sqr(n % 10)
		n /= 10
		if sum == 1 && n == 0 {
			return true
		}
		if n == 0 {
			n = sum
			sum = 0
		}
		tmp++
		if tmp > 100 {
			return false
		}

	}
	return false
}

func sqr(n int) int {
	return n * n
}

// _____________________
// 5 res

func isHappy4(n int) bool {
	for n != 1 {
		if n == 2 || n == 4 {
			return false
		}
		n = squareSum(n)
	}
	return true
}

func squareSum(n int) int {
	var r int
	for n != 0 {
		r += (n % 10) * (n % 10)
		n /= 10
	}
	return r
}
