package main

import (
	"fmt"
)

func majorityElement(nums []int) []int {
	major, count := make([]int, 2), make([]int, 2)

	max := 10000000001
	major[0], major[1] = max, max

	threshold := len(nums)/3 + len(nums)%2
	isLeft := true

	for i := 0; i < len(nums); i++ {
		if count[0] == 0 && nums[i] != major[1] {
			isLeft = true
			major[0] = nums[i]
			count[0] = 1
		} else if count[1] == 0 && nums[i] != major[0] {
			isLeft = false
			major[1] = nums[i]
			count[1] = 1
		} else if nums[i] == major[0] {
			isLeft = true
			count[0]++
		} else if nums[i] == major[1] {
			isLeft = false
			count[1]++
		} else if isLeft {
			count[1]--
		} else {
			count[0]--
		}
		fmt.Println(major, count, threshold)
	}

	rs := make([]int, 0)

	if count[0] > threshold && major[0] != max {
		rs = append(rs, major[0])

	}
	if count[1] > threshold && major[1] != max {
		rs = append(rs, major[1])

	}

	return rs
}

func main() {
	fmt.Println(majorityElement([]int{1, 2, 3, 4}))
}
