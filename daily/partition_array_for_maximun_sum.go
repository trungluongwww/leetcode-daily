package main

import (
	"fmt"
)

func maxSumAfterPartitioning(arr []int, k int) (total int) {

	for i := 0; i < len(arr)-k; i++ {
		max := arr[i]
		for j := i + 1; j < i+k; j++ {
			if arr[j] > max {
				max = arr[j]
			}
		}
		total += max
	}

	return total
}

func main() {
	fmt.Println(maxSumAfterPartitioning([]int{1, 15, 7, 9, 2, 5, 10}, 3))
}
