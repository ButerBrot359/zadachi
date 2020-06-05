package main

import "fmt"

func main() {
	a := make([]int, 5)
	a[0] = 2
	a[1] = 1
	a[2] = 5
	a[3] = 5
	a[4] = 1
	// fmt.Println(singleNumber(a))
	fmt.Println(singleNumber(a))
}

type Number struct {
	Data  int
	Count int
}

func singleNumber(nums []int) int {
	newMas := SortMasInt(nums)

	var a Number
	a.Data = newMas[0]
	a.Count = 1
	for i := 1; i < len(newMas); i++ {
		if newMas[i] != a.Data && a.Count%2 != 0 {
			return a.Data
		} else if newMas[i] != a.Data && a.Count%2 == 0 {
			a.Data = newMas[i]
			a.Count = 1
		} else if newMas[i] == a.Data {
			a.Count++
		}
	}
	return a.Data
}
func SortMasInt(nums []int) []int {
	a := nums
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[i] >= a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}
