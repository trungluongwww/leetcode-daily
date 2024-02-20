package main

func missingNumber(nums []int) int {
	sum := 0
	for i, num := range nums {
		sum += i - num
	}

	return sum + len(nums)
}
