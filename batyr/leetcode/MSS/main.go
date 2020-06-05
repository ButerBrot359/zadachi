package main

import "fmt"

func main() {
	// var a []int
	// // a = append(a, 3)
	// a = append(a, 2)
	// // a = append(a, 5)
	// a = append(a, 2)

	// fmt.Println(maxSubArray(a))

	a := 5
	b := 7

	fmt.Println(a, b)
}

// func maxSubArray(nums []int) int {
// 	return MaxSum(nums, len(nums))
// }
// func MaxSum(nums []int, n int) int {
// 	if n == 1 {
// 		return nums[0]
// 	}
// 	if n == 2 && nums[0]+nums[1] < Max(nums[0], nums[1]) {

// 		return Max(nums[0], nums[1])
// 	}
// 	neg := false
// 	for _, v := range nums {
// 		if v < 0 {
// 			neg = true
// 		} else {
// 			neg = false
// 			break
// 		}

// 	}
// 	if neg {
// 		return SortMaxInt(nums)
// 	}
// 	m := n / 2
// 	Left_MSS := MaxSum(nums, m)
// 	Right_MSS := MaxSum(nums[m:len(nums)], n-m)
// 	var Int_Min = 0
// 	Leftsum := Int_Min
// 	Rightsum := Int_Min
// 	sum := 0
// 	for i := m; i < n; i++ {
// 		sum += nums[i]
// 		Rightsum = Max(Rightsum, sum)
// 	}
// 	sum = 0
// 	for i := m - 1; i >= 0; i-- {
// 		sum += nums[i]
// 		Leftsum = Max(Leftsum, sum)
// 	}
// 	ans := Max(Left_MSS, Right_MSS)
// 	return Max(ans, Leftsum+Rightsum)
// }
// func Max(a, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }
// func SortMaxInt(nums []int) int {
// 	a := nums
// 	for i := 0; i < len(a); i++ {
// 		for j := i; j < len(a); j++ {
// 			if a[i] <= a[j] {
// 				a[i], a[j] = a[j], a[i]
// 			}
// 		}
// 	}
// 	return nums[0]
// }
