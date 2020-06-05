package main

import "fmt"

func main() {
	var a []int
	a = append(a, 6)
	a = append(a, 3)
	a = append(a, 2)
	a = append(a, 5)
	a = append(a, 8)
	a = append(a, 11)

	fmt.Print(lastStoneWeight(a))
}

func lastStoneWeight(stones []int) int {
	var mas []int
	if len(stones) == 1 {
		return stones[0]
	} else if len(stones) == 2 {
		return mod(stones[0] - stones[1])
	}
	SortMas(stones)
	if len(stones) >= 3 {
		mas = append(mas, stones[0]-stones[1])
		for i := 2; i < len(stones); i++ {
			mas = append(mas, stones[i])
		}
	}
	res := lastStoneWeight(mas)
	return res

}
func mod(num int) int {
	if num < 0 {
		num = -num
	}
	return num
}
func SortMas(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
