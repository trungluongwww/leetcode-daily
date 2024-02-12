package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	major, count := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			major = nums[i]
			count = 1
		} else if major == nums[i] {
			count++
		} else {
			count--
		}
	}
	return major
}

func main() {
	fmt.Println(majorityElement([]int{3, 3, 4}))
}
