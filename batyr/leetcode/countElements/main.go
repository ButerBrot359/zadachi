package main

func countElements(arr []int) int {

	if len(arr) < 1 && len(arr) > 1000 {
		return 0
	}
	for _, v := range arr {
		if v < 0 && v > 1000 {
			return 0
		}
	}
	count := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i]+1 == arr[j] {
				count++
				break
			}
		}
	}

	return count

}
