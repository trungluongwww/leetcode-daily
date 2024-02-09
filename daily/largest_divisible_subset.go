package main

import (
	"fmt"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	sort.Ints(nums)

	var (
		indexOfMaximum = 0
		maximum        = 0
		dp             = make([][]int, len(nums))
	)

	dp[0] = append(dp[0], nums[0])

	for i := 1; i < len(nums); i++ {
		dp[i] = append(dp[i], nums[i])

		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				if len(dp[j])+1 >= len(dp[i]) {
					dp[i] = append([]int{}, dp[j]...)
					dp[i] = append(dp[i], nums[i])
				}
			}
		}
		if len(dp[i]) > maximum {
			maximum = len(dp[i])
			indexOfMaximum = i
		}
	}

	return dp[indexOfMaximum]
}

func main() {
	fmt.Println(largestDivisibleSubset([]int{3, 4, 16, 8}))
}
