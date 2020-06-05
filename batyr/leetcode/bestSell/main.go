package main

import "fmt"

func main() {
	var a []int
	// a = append(a, 7)
	a = append(a, 1)
	a = append(a, 5)
	a = append(a, 6)
	a = append(a, 3)
	a = append(a, 8)

	fmt.Print(maxProfit(a))
}
func maxProfit(prices []int) int {
	goods := 0
	res := 0
	isBuy := false
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if !isBuy {
				if prices[j] <= prices[i] {
					i = j
				} else {
					goods = prices[i]
					isBuy = true
				}
			}
			if isBuy {
				if j == len(prices)-1 {
					res += prices[j] - goods
					return res
				} else if /*(goods < prices[j]) &&*/ prices[j] > prices[j+1] {
					res += prices[j] - goods
					isBuy = false
					i = j + 1
					continue
				}
			}
		}
	}
	return res
}

// func check(i, j int, prices []int) bool {

// }
