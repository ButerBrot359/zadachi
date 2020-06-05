package main

func maxSubArray1(nums []int) int {

	sum := nums[0]
	max := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] > sum && nums[i] > (sum+nums[i]) {
			sum = nums[i]
		} else {
			if i != 0 {
				sum = sum + nums[i]
			}
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

// ____________________________
func maxSubArray2(nums []int) int {
	var currsum int
	var maxsum = nums[0]
	for _, nb := range nums {

		currsum += nb
		if nb > currsum {
			currsum = nb
		}
		if currsum > maxsum {
			maxsum = currsum
		}
	}
	return maxsum
}

// _________________________________
func maxSubArray3(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// _________________________________
func maxSubArray4(nums []int) int {
	if len(nums) == 0 {
		return -2147483648
	}
	return DivideAndCon(nums, 0, len(nums)-1)
}

func DivideAndCon(nums []int, left, right int) int {

	if left == right {
		return nums[left]
	}
	mid := (left + right) / 2
	return MaxOfThree(DivideAndCon(nums, left, mid), DivideAndCon(nums, mid+1, right), CrossSubArray(nums, left, mid, right))
}

func CrossSubArray(nums []int, left, mid, right int) int {
	sum := 0
	left_sum := -2147483648

	for i := mid; i >= left; i-- {
		sum += nums[i]
		if sum > left_sum {
			left_sum = sum
		}
	}

	sum = 0
	right_sum := -2147483648
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		if sum > right_sum {
			right_sum = sum
		}
	}

	return left_sum + right_sum

}

func MaxOfThree(x, y, z int) int {
	max := x
	if y > max {
		max = y
	}

	if z > max {
		max = z
	}

	return max

}
