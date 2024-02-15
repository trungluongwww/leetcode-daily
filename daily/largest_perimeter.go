package main

import (
	"fmt"
	"sort"
)

func largestPerimeter(nums []int) int64 {
	sort.Ints(nums)

	sum := nums[0] + nums[1]

	var max int

	for i := 2; i < len(nums); i++ {
		if nums[i] < sum && nums[i]+sum > max {
			max = nums[i] + sum
		}

		sum += nums[i]
	}

	if max == 0 {
		return -1
	}

	return int64(max)
}

func main() {
	fmt.Println(largestPerimeter([]int{5, 5, 5}))

}
