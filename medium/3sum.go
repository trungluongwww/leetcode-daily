package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	rs := make([][]int, 0)

	sort.Ints(nums)

	if len(nums) < 3 {
		return rs
	}

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1

		for left < right {
			dif := nums[i] + nums[left] + nums[right]
			if dif == 0 {
				rs = append(rs, []int{nums[i], nums[left], nums[right]})
				for right > left && nums[right] == nums[right-1] {
					right--
				}

				for right > left && nums[left] == nums[left+1] {
					left++
				}
			} else if dif > 0 {
				right--
			} else {
				left++
			}

		}
	}

	return rs
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
