package main

func maxProfit1(prices []int) int {
	max := 0
	for i, j := 0, 1; j < len(prices); i, j = i+1, j+1 {
		if prices[j] > prices[i] {
			max += prices[j] - prices[i]
		}
	}
	return max
}

// _________
func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	profit := 0
	start := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] > prices[i+1] {
			profit += prices[i] - prices[start]
			start = i + 1
			continue
		}
	}
	if prices[len(prices)-1] > prices[start] {
		profit += prices[len(prices)-1] - prices[start]
	}
	return profit
}

// _______________
